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
import { Donation, DonationItem, AddDonationResponse } from "../../../types/donation";
import { ConfirmationModal } from "../../confirmation";


interface editingDonationItem {
    SKUName: string;
    quantity: number;
}

const AddDonationResponseModal =
    ({
         opened,
         setOpened,
         //refresh,
         takenItems,
     }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        //refresh: (search: string) => Promise<void>;
        takenItems: AddDonationResponse[];
}) => {
        return (
            <>
                <Modal
                    opened={opened}
                    onClose={() => {
                        //refresh("");
                        setOpened(false);
                    }}
                    title="Donation Items"
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
const DonationItemModal =
    ({
         opened,
         setOpened,
         refresh,
         items,
     }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        refresh: (search: string) => Promise<void>;
        items: DonationItem[];
}) => {
        return (
            <>
                <Modal
                    opened={opened}
                    onClose={() => {
                        refresh("");
                        setOpened(false);
                    }}
                    title="Donation Items"
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
                                <th>Total Credits</th>
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

const CreateDonationModal = ({
    opened,
    setOpened,
    refresh,
}: {
    opened: boolean;
    setOpened: Dispatch<SetStateAction<boolean>>;
    refresh: () => Promise<void>;
}) => {
    const [donationItems, setDonationItems] = useState<
        editingDonationItem[]
    >([]);

    const [quantity, setQuantity] = useState<number | "">('');

    const [suggestData, setSuggestData] = useState<any[]>([]);

    const [donorId, setDonorId] = useState<number>(0);
    const [items, setItems] = useState<Item[]>([]);
    const [itemSKUName, setItemSKUName] = useState('');
    const [showAddDonationResponseModal, setShowAddDonationResponseModal] = useState(false);

    const [disabled, setDisabled] = useState(false);

    const doSubmit = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donations/add`,
            {
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                method: "PUT",
                body: JSON.stringify({
                    donorId,
                    Items: donationItems,
                }),
            }
        );

        const data: {
            success: boolean;
            message: string;
        } = await response.json();

        if (data.success) {
            setShowAddDonationResponseModal(true);
            setOpened(false);
            refresh();
            showNotification({
                message: "Donation created successfully.",
                color: "green",
                title: "Donation created",
            });
            return
        }

        showNotification({
            message: data.message,
            color: "red",
            title: "Donation creation failed",
        });
        setDisabled(false);
    };

    const fetchDonors = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donor/list`,
            {
                credentials: "include",
            }
        );

        const data = await response.json();

        //console.log(data);



        if (data.success) {

            setSuggestData([])

            const donors = data.data

            let temp: any[]

            temp = []

            for (let i = 0; i < data.data.length; i++) {
                const name = donors[i].name
                const id = donors[i].ID
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
        if (opened) {
            setDisabled(false);
        }
    }, [opened]);
    
    useEffect(() => {
        fetchItems();
    },[]);
    
    
    useEffect(() => {
        fetchDonors();
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
                    setDisabled(false);
                    refresh();
                }}
                title="Create Donation"
            >
                <form
                    onSubmit={(e) => {
                        e.preventDefault();
                        setDisabled(true);
                        doSubmit();
                    }}
                >
                    
                    <Select
                        label="Donor"
                        required
                        data={suggestData}
                        onChange={(value) => {
                            setDonorId(Number(value));
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
                        <Space h="md" />
                        <Button
                            onClick={() => {
                                
                                const existingItem = donationItems.find(
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
                                    setDonationItems([
                                        {
                                            SKUName: itemSKUName,
                                            quantity,
                                        },
                                        ...donationItems,
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
                            {donationItems.map((item) => (
                                <tr key={item.SKUName}>
                                    <td>{item.SKUName}</td>
                                    <td>{item.quantity}</td>
                                    <td>
                                        <ActionIcon
                                            onClick={() => {
                                                // remove item from donation items
                                                setDonationItems(
                                                    donationItems.filter(
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
                        <Button type="submit" loading={disabled}>
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

const DonationComponent = ({
    donation,
    refresh,
}: {
    donation: Donation;
    refresh: () => Promise<void>;
}) => {

    const [donorBy, setDonorName] = useState<string>("");
    const [showItemModal, setShowItemModal] = useState(false);
    const [items, setItems] = useState<DonationItem[]>([]);
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
                    name: donorBy,
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

    // const getDonorName = async () => {
    //     const response = await fetch(
    //         `${process.env.REACT_APP_API_URL}/info/donor?id=${donation.donorId}`,
    //         {
    //             credentials: "include",
    //         }
    //     );

    //     const data: {
    //         success: boolean;
    //         data: string;
    //     } = await response.json();

    //     if (data.success) {
    //         setDonorName(data.data);
    //     }
    // };


    useEffect(() => {
        init();
    }, []);

    const init = async () => {
        //await getDonationItems();
        //await getPreparerUsername();
        //await getDonorName();
        await getItems();
    };

    const showDonationItems = async () => {

        getItems()
        setShowItemModal(true)

    }

    const getItems = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donations/items?id=${donation.ID}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: DonationItem[];
        } = await response.json();

        if (data.success) {
            console.log(data.data);
            setItems(data.data);
        }
    };

    const doDelete = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donations/delete?id=${donation.ID}`,
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
               
                <td>{donation.createdTime}</td>
                <td>{donation.signedBy}</td>
                <td>{donation.donorBy}</td>
                <td>{donation.totalValue}</td>
                <td>
                    <Group>
                        {/* <Tooltip label="Delete">
                            <ActionIcon variant="default" onClick={() => setShowConfirmationModal(true)}>
                                <Trash />
                            </ActionIcon>
                        </Tooltip> */}

                        <Tooltip label="List Items">
                            <ActionIcon
                                onClick={showDonationItems}
                                variant="default"
                            >
                                <ListDetails />
                            </ActionIcon>
                        </Tooltip>

                        {/* <Tooltip label = "Generate Invoice">
                            <ActionIcon
                                onClick={generateInvoice}
                                variant="default"
                            >
                                <FileInvoice />
                            </ActionIcon>
                        </Tooltip> */}
                    </Group>
                </td>
            </tr>
            <DonationItemModal
                opened={showItemModal}
                setOpened={setShowItemModal}
                refresh={showDonationItems}
                items={items}
            />
            <ConfirmationModal opened={showConfirmationModal} setOpened={setShowConfirmationModal} command={doDelete} message={"This action is not reversible. This will permanently delete the Donation beyond recovery."}/>
        </>
    );
};

export const DonationManager = () => {
    const [loading, setLoading] = useState<boolean>(false);

    const [currentPage, setCurrentPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);

    const [suggestData, setSuggestData] = useState<any[]>([]);

    const [donations, setDonations] = useState<Donation[]>([]);

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
        fetchDonors();
    }, []);

    useEffect(() => {
        fetchDonations();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        fetchDonations();
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

    const fetchDonors = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donor/list`,
            {
                credentials: "include",
            }
        );

        const data = await response.json();

        if (data.success) {

            setSuggestData([])

            const donors = data.data

            let temp: any[]

            temp = []

            for (let i = 0; i < data.data.length; i++) {
                const name = donors[i].name
                const id = donors[i].ID
                temp = [...temp, {value: id,label:name},]
            }

            setSuggestData(temp)


        }
    };

    const fetchDonations = async () => {
        setLoading(true);

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donations/list?page=${currentPage}&type=${typeFilter}&date=${dateFilter}&user=${userFilter}`,
            {
                credentials: "include",
            }
        );
        setLoading(false);

        const data: {
            success: boolean;
            message: string;
            data: {
                data: Donation[];
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setDonations(data.data.data);
            console.log(data.data.data);
            setTotalPages(data.data.totalPages);
            return;
        }
    };

    return (
        <>
            <Box sx={containerStyles}>
                <Group position="apart">
                    <h3>Donations</h3>{" "}
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
                        //label="Donor"
                        clearable
                        placeholder="Donor"
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
                            <th>Donated By</th>
                            <th>Total Credits</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {donations.map((donation) => (
                            <DonationComponent
                                refresh={fetchDonations}
                                key={donation.ID}
                                donation={donation}
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
                    <Button onClick={fetchDonations}>
                        Refresh
                    </Button>
                </Group>
            </Box>
            <CreateDonationModal
                opened={showCreationModal}
                setOpened={setShowCreationModal}
                refresh={fetchDonations}
            />
        </>
    );
};

export default DonationManager;

