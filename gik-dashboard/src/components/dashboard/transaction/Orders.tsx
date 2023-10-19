import {
    ActionIcon,
    Autocomplete,
    Box,
    Button,
    Center,
    Group,
    Modal,
    Pagination,
    Select,
    Space,
    Table,
    NumberInput,
    Tooltip,
} from "@mantine/core";
import { DatePickerInput } from "@mantine/dates";
import { showNotification } from "@mantine/notifications";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { CirclePlus, Trash, ListDetails, FileInvoice } from "tabler-icons-react";
import { containerStyles } from "../../../styles/container";
import { Item } from "../../../types/item";
import { Order, OrderItem, AddOrderResponse } from "../../../types/order";
import { ConfirmationModal } from "../../confirmation";


interface editingOrderItem {
    SKUName: string;
    quantity: number;
}

const AddOrderResponseModal =
    ({
         opened,
         setOpened,
         //refresh,
         takenItems,
     }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        //refresh: (search: string) => Promise<void>;
        takenItems: AddOrderResponse[];
}) => {
        return (
            <>
                <Modal
                    opened={opened}
                    onClose={() => {
                        //refresh("");
                        setOpened(false);
                    }}
                    title="Order Items"
                    size="50%"
                >
                    <Box sx={containerStyles}>
                        <Space h="md" />
                        <Table>
                            <thead>
                            <tr>
                                <th>Item</th>
                                <th>Location</th>
                                <th>Taken Quantity</th>
                            </tr>
                            </thead>
                            <tbody>
                            {takenItems.map((takenItem) => (
                                <tr key={takenItem.ID}>
                                    <td>{takenItem.itemSKUName}</td>
                                    <td>{takenItem.locationName}</td>
                                    <td>{takenItem.quantity}</td>
                                </tr>
                            ))}
                            </tbody>
                        </Table>

                    </Box>
                </Modal>
            </>
    );
};
const OrderItemModal =
    ({
         opened,
         setOpened,
         refresh,
         items,
     }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        refresh: (search: string) => Promise<void>;
        items: OrderItem[];
}) => {
        return (
            <>
                <Modal
                    opened={opened}
                    onClose={() => {
                        refresh("");
                        setOpened(false);
                    }}
                    title="Order Items"
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
                                <th>Total Cost</th>
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
                                    <td>{item.totalCost}</td>
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
    const [orderItems, setOrderItems] = useState<
        editingOrderItem[]
    >([]);

    const [quantity, setQuantity] = useState<number | "">('');

    const [suggestData, setSuggestData] = useState<any[]>([]);

    const [clientId, setClientId] = useState<number>(0);
    const [items, setItems] = useState<Item[]>([]);
    const [itemSKUName, setItemSKUName] = useState('');
    const [showAddOrderResponseModal, setShowAddOrderResponseModal] = useState(false);

    const doSubmit = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/orders/add`,
            {
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                method: "PUT",
                body: JSON.stringify({
                    clientId,
                    Items: orderItems,
                }),
            }
        );

        const data: {
            success: boolean;
            message: string;
            data : AddOrderResponse[];
        } = await response.json();

        if (data.success) {
            setShowAddOrderResponseModal(true);
            setOpened(false);
            refresh();
            // showNotification({
            //     message: "Order created successfully.",
            //     color: "green",
            //     title: "Order created",
            // });
            let addOrderMsg = "";
            let idx = 0;
            //
            addOrderMsg = addOrderMsg + "List and amount of item to be taken:\n"
            addOrderMsg += "---------------\n";
            //
            for (idx=0; idx<data.data.length; idx++){
                const tmpItem = data.data[idx];
                addOrderMsg = addOrderMsg + tmpItem.ID.toString() + ". Take " + tmpItem.quantity.toString() +" of "+ tmpItem.itemSKUName + " from location " + tmpItem.locationName + "\n"
            }
            alert(addOrderMsg);
            setOrderItems([]);
            console.log(data.data);
            return (
            <AddOrderResponseModal
                opened={showAddOrderResponseModal}
                setOpened={setShowAddOrderResponseModal}
                //refresh={CreateOrderModal}
                takenItems={data.data}
            />
            );
           
        }

        showNotification({
            message: data.message,
            color: "red",
            title: "Order creation failed",
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

    const fetchItems = async () => {

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/list`,
            {
                credentials: "include",
            }
        );


        const data: {
            success: boolean;
            message: string;
            data: {
                data: Item[];
                total: number;
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            console.log("Success loading item data");
            //console.log(data.data.data)
            setItems(data.data.data);
            console.log(data.data.data);
        }
    };
    
    
    useEffect(() => {
        fetchItems();
    },[]);
    
    
    useEffect(() => {
        fetchClients();
    }, []);

    const lstItemSKUName : string[] = [];
    for(let idx = 0; idx<items.length; idx++){
        lstItemSKUName.push(items[idx].sku + " : " + items[idx].name);
    }

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
                    {/* <Space h="md" />
                    <Group
                        grow
                        sx={{
                            alignItems: "flex-end",
                        }}
                    > */}
                        <Autocomplete
                            label="Items"
                            placeholder="Name or SKU (type any word to search)"
                            data={lstItemSKUName}
                            value={itemSKUName}
                            onChange={setItemSKUName}
                            
                        />
                        <NumberInput
                        label= "Quantity"
                        placeholder= "10"
                        min = {0}
                        value = {quantity}
                        onChange={setQuantity}
                        />
                        <Button
                            onClick={() => {
                                
                                const existingItem = orderItems.find(
                                    (item) => item.SKUName === itemSKUName
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

                                if(quantity != "") {
                                    setOrderItems([
                                        {
                                            SKUName: itemSKUName,
                                            quantity,
                                        },
                                        ...orderItems,
                                    ]);
                                }
                                setItemSKUName("");
                                setQuantity(0);
                            }}
                        >
                            Add
                        </Button>
                    {/* </Group> */}
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
                            {orderItems.map((item) => (
                                <tr key={item.SKUName}>
                                    <td>{item.SKUName}</td>
                                    <td>{item.quantity}</td>
                                    <td>
                                        <ActionIcon
                                            onClick={() => {
                                                // remove item from order items
                                                setOrderItems(
                                                    orderItems.filter(
                                                        (i) => i.SKUName !== item.SKUName
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

const OrderComponent = ({
    order,
    refresh,
}: {
    order: Order;
    refresh: () => Promise<void>;
}) => {

    const [clientName, setClientName] = useState<string>("");
    const [showItemModal, setShowItemModal] = useState(false);
    const [items, setItems] = useState<OrderItem[]>([]);
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
            `${process.env.REACT_APP_API_URL}/info/client?id=${order.clientId}`,
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

    // const getPreparerUsername = async () => {
    //     const response = await fetch(
    //         `${process.env.REACT_APP_API_URL}/info/username?id=${order.signerId}`,
    //         {
    //             credentials: "include",
    //         }
    //     );

    //     const data: {
    //         success: boolean;
    //         data: string;
    //     } = await response.json();

    //     if (data.success) {
    //         setPreparerUsername(data.data);
    //     }
    // };

    useEffect(() => {
        init();
    }, []);

    const init = async () => {
        //await getOrderItems();
        //await getPreparerUsername();
        await getClientName();
        await getItems();
    };

    const showOrderItems = async () => {

        getItems()
        setShowItemModal(true)

    }

    const getItems = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/orders/items?id=${order.ID}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: OrderItem[];
        } = await response.json();

        if (data.success) {
            console.log(data.data);
            setItems(data.data);
        }
    };

    const doDelete = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/orders/delete?id=${order.ID}`,
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
               
                <td>{order.createdTime}</td>
                <td>{order.signedBy}</td>
                <td>{order.clientName}</td>
                <td>{order.totalCost}</td>
                <td>
                    <Group>
                        <Tooltip label="Delete">
                            <ActionIcon variant="default" onClick={() => setShowConfirmationModal(true)}>
                                <Trash />
                            </ActionIcon>
                        </Tooltip>

                        <Tooltip label="List Items">
                            <ActionIcon
                                onClick={showOrderItems}
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
            <OrderItemModal
                opened={showItemModal}
                setOpened={setShowItemModal}
                refresh={showOrderItems}
                items={items}
            />
            <ConfirmationModal opened={showConfirmationModal} setOpened={setShowConfirmationModal} command={doDelete} message={"This action is not reversible. This will permanently delete the Order beyond recovery."}/>
        </>
    );
};

export const OrderManager = () => {
    const [loading, setLoading] = useState<boolean>(false);

    const [currentPage, setCurrentPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);

    const [suggestData, setSuggestData] = useState<any[]>([]);

    const [orders, setOrders] = useState<Order[]>([]);

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
        fetchOrders();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        fetchOrders();
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

    const fetchOrders = async () => {
        setLoading(true);

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/orders/list?page=${currentPage}&type=${typeFilter}&date=${dateFilter}&user=${userFilter}`,
            {
                credentials: "include",
            }
        );
        setLoading(false);

        const data: {
            success: boolean;
            message: string;
            data: {
                data: Order[];
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setOrders(data.data.data);
            console.log(data.data.data);
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
                            <th>Time</th>
                            <th>Signed By</th>
                            <th>Client</th>
                            <th>Total Cost</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {orders.map((order) => (
                            <OrderComponent
                                refresh={fetchOrders}
                                key={order.ID}
                                order={order}
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
                    <Button onClick={fetchOrders}>
                        Refresh
                    </Button>
                </Group>
            </Box>
            <CreateOrderModal
                opened={showCreationModal}
                setOpened={setShowCreationModal}
                refresh={fetchOrders}
            />
        </>
    );
};

export default OrderManager;

