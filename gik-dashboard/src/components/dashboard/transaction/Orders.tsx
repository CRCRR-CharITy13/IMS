import {
    ActionIcon,
    Box,
    Button,
    Center,
    Group,
    Modal,
    Pagination,
    Select,
    Space,
    Table,
    TextInput,
    Tooltip,
} from "@mantine/core";
import { DatePickerInput } from "@mantine/dates";
import { showNotification } from "@mantine/notifications";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { CirclePlus, Trash, ListDetails, FileInvoice } from "tabler-icons-react";
import { containerStyles } from "../../../styles/container";
import { Client } from "../../../types/client";
import { Transaction, TransactionItem } from "../../../types/transaction";
import { ConfirmationModal } from "../../confirmation";


interface editingTransactionItem {
    id: number;
    quantity: number;
}


const TransactionItemModal =
    ({
         opened,
         setOpened,
         refresh,
         items,
     }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        refresh: (search: string) => Promise<void>;
        items: TransactionItem[];
}) => {
        //const [items, setItems] = useState<TransactionItem[]>([]);
        return (
            <>
                <Modal
                    opened={opened}
                    onClose={() => {
                        refresh("");
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
                                <th>Quantity</th>
                                <th>Total Value</th>
                            </tr>
                            </thead>
                            <tbody>
                            {items.map((item) => (
                                <tr key={item.ID}>
                                    <td>{item.name}</td>
                                    <td>{item.sku}</td>
                                    <td>{item.size}</td>
                                    <td>{item.price}</td>
                                    <td>{item.quantity}</td>
                                    <td>{item.totalValue}</td>
                                </tr>
                            ))}
                            </tbody>
                        </Table>

                    </Box>
                </Modal>
            </>
    );
};

const CreateOrderModal = ({
    opened,
    setOpened,
    refresh,
}: {
    opened: boolean;
    setOpened: Dispatch<SetStateAction<boolean>>;
    refresh: () => Promise<void>;
}) => {
    const [transactionItems, setTransactionItems] = useState<
        editingTransactionItem[]
    >([]);

    const [productId, setProductId] = useState<number>(0);
    const [quantity, setQuantity] = useState<number>(0);

    const [suggestData, setSuggestData] = useState<any[]>([]);

    const [clientId, setClientId] = useState<number>(0);


    const doSubmit = async () => {
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
            setTransactionItems([]);
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

            const clients = data.data

            let temp: any[]

            temp = []

            for (let i = 0; i < data.data.length; i++) {
                const name = clients[i].name
                const id = clients[i].ID
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
                    refresh();
                }}
                title="Create Order"
            >
                <form
                    onSubmit={(e) => {
                        e.preventDefault();
                        doSubmit();
                    }}
                >
                    
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
                    <Group
                        grow
                        sx={{
                            alignItems: "flex-end",
                        }}
                    >
                        <TextInput
                            label="Product ID"
                            required
                            placeholder="19281"
                            type="number"
                            value={productId}
                            onChange={(e) =>
                                setProductId(Number(e.target.value))
                            }
                        />
                        <TextInput
                            label="Quantity"
                            required
                            placeholder="6"
                            type="number"
                            value={quantity}
                            onChange={(e) =>
                                setQuantity(Number(e.target.value))
                            }
                        />
                        <Button
                            onClick={() => {
                                // check if product exists in transaction items already
                                const existingItem = transactionItems.find(
                                    (item) => item.id === Number(productId)
                                );

                                if (existingItem) {
                                    showNotification({
                                        color: "red",
                                        title: "Item already exists",
                                        message:
                                            "Please remove the item first.",
                                    });

                                    return;
                                }

                                setTransactionItems([
                                    {
                                        id: Number(productId),
                                        quantity,
                                    },
                                    ...transactionItems,
                                ]);
                                setProductId(0);
                                setQuantity(0);
                            }}
                        >
                            Add
                        </Button>
                    </Group>
                    <Space h="md" />
                    <Table>
                        <thead>
                            <tr>
                                <th>Product ID</th>
                                <th>Quantity</th>
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {transactionItems.map((item) => (
                                <tr key={item.id}>
                                    <td>{item.id}</td>
                                    <td>{item.quantity}</td>
                                    <td>
                                        <ActionIcon
                                            onClick={() => {
                                                // remove item from transaction items
                                                setTransactionItems(
                                                    transactionItems.filter(
                                                        (i) => i.id !== item.id
                                                    )
                                                );
                                            }}
                                            variant="default"
                                        >
                                            <Trash />
                                        </ActionIcon>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </Table>
                    <Space h="md" />
                    <Group position="right">
                        <Button type="submit">
                            Submit
                        </Button>
                    </Group>
                </form>
            </Modal>
        </>
    );
};

interface invoiceItem {
    name: string;
    size: string;
    sku: string;
    price: number;
    quantity: number;
}

const TransactionComponent = ({
    transaction,
    refresh,
}: {
    transaction: Transaction;
    refresh: () => Promise<void>;
}) => {

    const [preparerUsername, setPreparerUsername] = useState<string>("");
    const [clientName, setClientName] = useState<string>("");


    const [showItemModal, setShowItemModal] = useState(false);
    const [items, setItems] = useState<TransactionItem[]>([]);

    const [showConfirmationModal, setShowConfirmationModal] =
        useState<boolean>(false);


    const generateInvoice = async () => {

        await getItems()

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/invoice/generate`,
            {
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                method: "POST",
                body: JSON.stringify({
                    name: clientName,
                    address: "customerAddress",
                    data: items,
                }),
            }
        );

        const data = await response.arrayBuffer();

        const blob = new Blob([data], { type: "application/pdf" });

        const url = URL.createObjectURL(blob);

        window.open(url);
    };

    const getClientName = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/info/client?id=${transaction.clientId}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: string;
        } = await response.json();

        if (data.success) {
            setClientName(data.data);
        }
    };

    const getPreparerUsername = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/info/username?id=${transaction.signerId}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: string;
        } = await response.json();

        if (data.success) {
            setPreparerUsername(data.data);
        }
    };

    useEffect(() => {
        init();
    }, []);

    const init = async () => {
        //await getTransactionItems();
        await getPreparerUsername();
        await getClientName();
        await getItems()
    };

    const showTransactionItems = async () => {

        getItems()
        setShowItemModal(true)

    }

    const getItems = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/transaction/items?id=${transaction.ID}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: TransactionItem[];
        } = await response.json();

        if (data.success) {
            setItems(data.data);
        }
    };

    const doDelete = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/transaction/delete?id=${transaction.ID}`,
            {
                credentials: "include",
                method: "DELETE",
            }
        );

        const data: {
            success: boolean;
        } = await response.json();

        if (data.success) {
            refresh();
        }
    };

    return (
        <>
            <tr>
                <td>{transaction.ID}</td>
                <td>
                    {new Date(transaction.timestamp * 1000).toLocaleString()}
                </td>
                <td>{transaction.type ? "Import" : "Export"}</td>
                <td>{preparerUsername}</td>
                <td>{clientName}</td>
                <td>{transaction.totalQuantity}</td>
                <td>
                    <Group>
                        <Tooltip label="Delete">
                            <ActionIcon variant="default" onClick={() => setShowConfirmationModal(true)}>
                                <Trash />
                            </ActionIcon>
                        </Tooltip>

                        <Tooltip label="List Items">
                            <ActionIcon
                                onClick={showTransactionItems}
                                variant="default"
                            >
                                <ListDetails />
                            </ActionIcon>
                        </Tooltip>

                        <Tooltip label = "Generate Invoice">
                            <ActionIcon
                                onClick={generateInvoice}
                                variant="default"
                            >
                                <FileInvoice />
                            </ActionIcon>
                        </Tooltip>
                    </Group>
                </td>
            </tr>
            <TransactionItemModal
                opened={showItemModal}
                setOpened={setShowItemModal}
                refresh={showTransactionItems}
                items={items}
            />
            <ConfirmationModal opened={showConfirmationModal} setOpened={setShowConfirmationModal} command={doDelete} message={"This action is not reversible. This will permanently delete the Transaction beyond recovery."}/>
        </>
    );
};

export const OrderManager = () => {
    const [loading, setLoading] = useState<boolean>(false);

    const [currentPage, setCurrentPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);

    const [suggestData, setSuggestData] = useState<any[]>([]);

    const [transactions, setTransactions] = useState<Transaction[]>([]);

    const [showCreationModal, setShowCreationModal] = useState(false);

    const [dateFilter, setDateFilter] = useState<
        [number | null, number | null] | undefined
        >();
    const [dateFilterEditing, setDateFilterEditing] = useState<
        [Date | null, Date | null] | undefined
        >();

    const [userFilter, setUserFilter] = useState<number>(0);
    const [userFilterEditing, setUserFilterEditing] = useState<number>(0);

    const [typeFilter, setTypeFilter] = useState<string>("");
    const [typeFilterEditing, setTypeFilterEditing] = useState<string>("");

    useEffect(() => {
        fetchClients();
    }, []);

    useEffect(() => {
        fetchTransactions();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        fetchTransactions();
    }, [dateFilter, userFilter, typeFilter]);

    const doFilter = async () => {
        // console.log("Filter Button Pressed");
        setTypeFilter(typeFilterEditing);
        if (dateFilterEditing != null && dateFilterEditing[0] != null && dateFilterEditing[1] != null) {
            setDateFilter([dateFilterEditing?.[0].getTime()/1000, dateFilterEditing?.[1].getTime()/1000]);
        } else {
            setDateFilter([null, null]);
        }
        setUserFilter(userFilterEditing);
    };

    const fetchClients = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/client/list`,
            {
                credentials: "include",
            }
        );

        const data = await response.json();

        if (data.success) {

            setSuggestData([])

            const clients = data.data

            let temp: any[]

            temp = []

            for (let i = 0; i < data.data.length; i++) {
                const name = clients[i].name
                const id = clients[i].ID
                temp = [...temp, {value: id,label:name},]
            }

            setSuggestData(temp)


        }
    };

    const fetchTransactions = async () => {
        setLoading(true);

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/transaction/list?page=${currentPage}&type=${typeFilter}&date=${dateFilter}&user=${userFilter}`,
            {
                credentials: "include",
            }
        );
        setLoading(false);

        const data: {
            success: boolean;
            message: string;
            data: {
                data: Transaction[];
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setTransactions(data.data.data);
            setTotalPages(data.data.totalPages);
            return;
        }
    };

    return (
        <>
            <Box sx={containerStyles}>
                <Group position="apart">
                    <h3>Orders</h3>{" "}
                    <ActionIcon
                        sx={{
                            height: "4rem",
                            width: "4rem",
                        }}
                        onClick={() => setShowCreationModal(true)}
                    >
                        <CirclePlus />
                    </ActionIcon>
                </Group>
                <Space h="md" />
                <Group>
                    <DatePickerInput
                        type="range"
                        placeholder="Pick Date Range"
                        onChange={setDateFilterEditing}
                    />
                    <Select
                        //label="Client"
                        placeholder="Client"
                        data={suggestData}
                        onChange={(value) => {
                            setUserFilterEditing(Number(value));
                            // console.log("onChange: "+userFilterEditing);
                        }}
                    />
                  
                    <Button onClick={doFilter}>
                        Filter
                    </Button>
                </Group>
                <Table>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Timestamp</th>
                            <th>Type</th>
                            <th>Preparer</th>
                            <th>Customer</th>
                            <th>Item Quantity</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {transactions.map((transaction) => (
                            <TransactionComponent
                                refresh={fetchTransactions}
                                key={transaction.ID}
                                transaction={transaction}
                            />
                        ))}
                    </tbody>
                </Table>
                <Space h="md" />
                <Center>
                    <Pagination
                        value={currentPage}
                        total={totalPages}
                        onChange={setCurrentPage}
                    />
                </Center>
                <Group position="right">
                    <Button onClick={fetchTransactions}>
                        Refresh
                    </Button>
                </Group>
            </Box>
            <CreateOrderModal
                opened={showCreationModal}
                setOpened={setShowCreationModal}
                refresh={fetchTransactions}
            />
        </>
    );
};

export default OrderManager;

