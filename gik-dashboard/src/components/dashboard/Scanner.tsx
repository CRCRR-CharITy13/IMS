import {
    Box,
    Modal,
    Container,
    Button,
    Group,
    TextInput,
    NumberInput, NumberInputHandlers,
    Space, Table, ActionIcon, Select,
} from "@mantine/core";
import { useForceUpdate } from '@mantine/hooks';
import { showNotification } from "@mantine/notifications";
import { Dispatch, SetStateAction, useEffect, useState, useRef } from "react";

import { QrReader } from "react-qr-reader";
import { useNavigate } from "react-router-dom";
import { Item } from "../../types/item";
import {TransactionItem} from "../../types/transaction";
import {containerStyles} from "../../styles/container";
import {CirclePlus, Trash} from "tabler-icons-react";

interface editingTransactionItem {
    id: number;
    quantity: number;
}

const CreateTransactionModal = (
    {
        opened,
        setOpened,
        refresh,
        transactionItems,
    }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        refresh: () => Promise<void>;
        transactionItems: editingTransactionItem[]
}) => {

    const [suggestData, setSuggestData] = useState<any[]>([]);

    const [clientId, setClientId] = useState<number>(0);

    const [transactionType, setTransactionType] = useState<boolean>(false);

    const doSubmit = async () => {
        console.log(transactionItems)
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/transaction/add`,
            {
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                method: "PUT",
                body: JSON.stringify({
                    clientId,
                    type: transactionType,
                    products: transactionItems,
                }),
            }
        );

        const data: {
            success: boolean;
            message: string;
        } = await response.json();

        if (data.success) {
            setOpened(false);
            refresh();
            showNotification({
                message: "Transaction created successfully.",
                color: "green",
                title: "Transaction created",
            });
            // setSuggestData([]);
            return;
        }

        showNotification({
            message: data.message,
            color: "red",
            title: "Transaction creation failed",
        });
    };

    const fetchClients = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/client/list`,
            {
                credentials: "include",
            }
        );

        const data = await response.json();

        //console.log(data);



        if (data.success) {

            setSuggestData([])

            let clients = data.data

            let temp: any[]

            temp = []

            for (let i = 0; i < data.data.length; i++) {
                let name = clients[i].name
                let id = clients[i].ID
                temp = [...temp, {value: id,label:name},]
            }

            setSuggestData(temp)


        }
    };

    useEffect(() => {
        fetchClients();
    }, []);

    return (
        <>
            <Modal
                opened={opened}
                onClose={() => {
                    setOpened(false);
                }}
                title="Create Transaction"
            >
                <form
                    onSubmit={(e) => {
                        e.preventDefault();
                        doSubmit();
                    }}
                >
                    <Select
                        required
                        label="Transaction Type"
                        data={[
                            {
                                value: "true",
                                label: "Import",
                            },
                            {
                                value: "false",
                                label: "Export",
                            },
                        ]}
                        value={transactionType.toString()}
                        onChange={(value) => {
                            setTransactionType(Boolean(value));
                        }}
                    />
                    <Space h="md" />
                    <Select
                        label="Client"
                        required
                        data={suggestData}
                        onChange={(value) => {
                            setClientId(Number(value));
                        }}
                    />
                    <Space h="md" />
                    <Group position="right">
                        <Button color="green" type="submit">
                            Submit
                        </Button>
                    </Group>
                </form>
            </Modal>
        </>
    );
};

const PermissionsModal = ({ opened }: { opened: boolean }) => {
    const navigate = useNavigate();

    return (
        <>
            <Modal
                onClose={() => {}}
                opened={opened}
                closeOnClickOutside={false}
                closeOnEscape={false}
                withCloseButton={false}
                title="Permissions Needed"
            >
                <h1>This application requires the camera.</h1>
                <p>Please activate the camera permission to enable scanning.</p>
                <Button
                    color="red"
                    onClick={() =>
                        navigate("/dashboard/analytics", { replace: true })
                    }
                >
                    Leave
                </Button>
            </Modal>
        </>
    );
};



const LocationItemRow = ({
                             i,
                             setStock,
                             stock,
                         }: {
    i: number;
    setStock: Dispatch<SetStateAction<any[]>>;
    stock: any[];
}) => {
    const handlers = useRef<NumberInputHandlers>();
    return (
        <Group>

            {/*@ts-ignore*/}
            <ActionIcon size={42} variant="default" onClick={() => handlers.current.decrement()}>
                –
            </ActionIcon>

            <NumberInput
                hideControls
                value={stock[i][1]}
                onChange={(val) => {
                    let temp: any[]

                    temp = stock

                    if (val == undefined) {
                        val = 0
                    }


                    temp[i][1] = val

                    setStock(temp)
                }}
                handlersRef={handlers}

                min={0}
                step={1}
                styles={{ input: { width: 54, textAlign: 'center' } }}
            />

            {/*@ts-ignore*/}
            <ActionIcon size={42} variant="default" onClick={() => handlers.current.increment()}>
                +
            </ActionIcon>
        </Group>
    )
}

const TransactionItemRow = (
    ({item, id, doDelete}: {item: any[], id:number, doDelete: (id: number) => void}) => {

        const [itemInfo, setItemInfo] = useState<Item>()


        const getItemInfo = async () => {
            if (itemInfo == undefined) {
                const response = await fetch(
                    `${process.env.REACT_APP_API_URL}/items/lookup?id=${item[0]}`,
                    {
                        credentials: "include",
                    }
                );

                const data: {
                    success: boolean;
                    data: Item;
                } = await response.json();

                if (data.success) {
                    await setItemInfo(data.data)
                }

            }

        }

        getItemInfo()

        if (itemInfo == undefined) {return (<></>)}
        return (
            <tr key={itemInfo.ID}>
                <td>{itemInfo.name}</td>
                <td>{itemInfo.sku}</td>
                <td>{itemInfo.size}</td>
                <td>{itemInfo.price}</td>
                <td>{item[1]}</td>
                <td>
                    <ActionIcon variant="default" onClick={() => {
                        doDelete(id)
                    }}>
                        <Trash />
                    </ActionIcon>
                </td>
            </tr>
        )


    }
)


const LocationItemModal =
    ({
         opened,
         setOpened,
         setItems,
         items,
     }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        setItems: (items: any[]) => void;
        items: Item[];
    }) => {
        const [stock, setStock] = useState<any[]>([]);
        const handlers = useRef<NumberInputHandlers>();
        const [testStock, setTestStock] = useState([0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0])

        useEffect(() => {
            init();
        }, [opened]);

        const init = async () => {
            let temp: any[]

            temp = []

            for (let i = 0; i < items.length; i++) {
                temp = [...temp, [items[i].ID, 0]]
            }
            setStock(temp)
        }
        return (
            <>
                <Modal
                    opened={opened}
                    onClose={() => {
                        setItems(stock);
                        setOpened(false);
                    }}
                    title="Transaction Items"
                    size="50%"
                >
                    <Box sx={containerStyles}>
                        <Space h="md" />
                        <Table>
                            <thead>
                            <tr>
                                <th>Name</th>
                                <th>SKU</th>
                                <th>Size</th>
                                <th>Price</th>
                                <th>Stock</th>
                                <th>Quantity</th>
                            </tr>
                            </thead>
                            <tbody>
                            {items.map((item, i) => (

                                <tr key={item.ID}>
                                    <td>{item.name}</td>
                                    <td>{item.sku}</td>
                                    <td>{item.size}</td>
                                    <td>{item.price}</td>
                                    <td>{item.quantity}</td>
                                    <td>
                                        <LocationItemRow
                                            i={i}
                                            setStock={setStock}
                                            stock={stock}
                                        />
                                    </td>
                                </tr>
                            ))}
                            </tbody>
                        </Table>

                    </Box>
                </Modal>
            </>
        );
    };


const Scanner = () => {
    const [results, setResults] = useState<string[]>([]);

    const [modalOpened, setModalOpened] = useState<boolean>(false);
    const [completeTransaction, setCompleteTransaction] = useState<boolean>(false);

    const [permissionsAllowed, setPermissionsAllowed] = useState<boolean>(true);

    const [transactionItems, setTransactionItems] = useState<any[]>([])
    const [transactionItemsPost, setTransactionItemsPost] = useState<editingTransactionItem[]>([])

    const [locationItems, setLocationItems] = useState<Item[]>([])

    const forceUpdate = useForceUpdate();

    useEffect(() => {
        checkPermission();
    }, []);

    useEffect(() => {
        if (!transactionItems) return;

    }, [transactionItems]);

    const resetItems = async () => {
        setTransactionItems([])
        forceUpdate()
    }

    const removeItem = (id: number) => {
        let tempTransaction = transactionItems
        tempTransaction.splice(id, 1)
        setTransactionItems(tempTransaction)
        forceUpdate()
    }

    const setItemHandler = (items: any[]) => {
        let tempTransaction = transactionItems

        let id: number;

        for (let i = 0; i < items.length; i++) {

            id = -1
            if (items[i][1] != 0) {
                for (let k = 0; k < tempTransaction.length; k++) {

                    console.log("_______________________")
                    console.log(tempTransaction[k][0])
                    console.log(items[i][0])



                    if (tempTransaction[k][0] == items[i][0] ) {
                        id = k
                        break
                    }

                }
                if (id == -1) {
                    tempTransaction = [...tempTransaction, items[i]]
                    id = tempTransaction.length - 1
                }
                else {
                    tempTransaction[id][1] += items[id][1]
                    if (tempTransaction[id][1] == 0) {
                        tempTransaction.splice(id, 1)
                    }
                }

            }
        }
        setTransactionItems(tempTransaction)

    }

    const checkPermission = async () => {
        // let allowed = false;

        // try {
        //     allowed = !!(await navigator.mediaDevices.getUserMedia({
        //         video: true,
        //     }));
        // } catch (err) {
        //     /* handle the error */
        //     allowed = false;
        // }

        setPermissionsAllowed(true);
    };

    const getItems = async (location: string) => {
        const splitResult = location.split("");

        const letter = splitResult[splitResult.length - 1];

        // get name which is everything except last character
        const name = location.slice(0, -1);

        const responseLocation = await fetch(
            `${process.env.REACT_APP_API_URL}/location/list/sku?name=${name}&letter=${letter}`,
            {
                credentials: "include",
            }
        );

        const dataLocation: {
            success: boolean;
            sku: string;
        } = await responseLocation.json();


        let sku = dataLocation.sku

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/list?sku=${sku}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: {
                data:        Item[],
                total:       number,
                currentPage: number,
                totalPages:  number,
            };
        } = await response.json();

        if (data.success) {
            await setLocationItems(data.data.data);
            setModalOpened(true)
        }
    };


    return (
        <>
            <Container>
                <h1>Item Scanner</h1>

                <h2>
                    Simply point your device at a QR code and wait for the
                    results to appear.
                </h2>
                <QrReader
                    onResult={(result, err) => {
                        if (!!result){
                            // @ts-ignore
                            if (!result?.text || results.includes(result?.text))
                                return;
                            // @ts-ignore
                            getItems(result?.text);
                        }

                    }}
                    constraints={{ facingMode: "environment" }}
                    scanDelay={500}
                />
            </Container>

            <Box sx={containerStyles}>
                <Group position="left">
                    <h3>Transactions Items</h3>
                </Group>
                <Table>
                    <thead>
                    <tr>
                        <th>Name</th>
                        <th>SKU</th>
                        <th>Size</th>
                        <th>Price</th>
                        <th>Quantity</th>
                        <th>Delete</th>
                    </tr>
                    </thead>
                    <tbody>
                    {transactionItems.map((item, i) => (

                        <TransactionItemRow
                            item={item}
                            id={i}
                            doDelete={removeItem}
                        />

                    ))}
                    </tbody>
                </Table>
                <Space h="md" />
                <Group position="right">
                    <Button
                        color="green"
                        onClick={() => {
                            if (transactionItems.length == 0) {return}
                            let temp: editingTransactionItem[]
                            temp = []

                            for (let i = 0; i < transactionItems.length; i++){
                                temp = [...temp, {
                                    id: Number(transactionItems[i][0]),
                                    quantity: transactionItems[i][1],
                                }]
                            }
                            setTransactionItemsPost(temp)

                            setCompleteTransaction(true)
                        }}
                    >
                        Complete
                    </Button>
                </Group>
            </Box>
            <LocationItemModal
                opened={modalOpened}
                setOpened={setModalOpened}
                setItems={setItemHandler}
                items={locationItems}
            />
            <CreateTransactionModal
                opened={completeTransaction}
                setOpened={setCompleteTransaction}
                refresh={resetItems}
                transactionItems={transactionItemsPost}
            />
            <PermissionsModal opened={!permissionsAllowed} />
        </>
    );
};

export default Scanner;
