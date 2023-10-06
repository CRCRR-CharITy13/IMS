import {
    Group,
    TextInput,
    Button,
    Center,
    Pagination,
    Space,
    Loader,
    Box,
    Table,
    Modal,
    ActionIcon,
    HoverCard,
    MultiSelect,
    Menu,
    Text,
    NumberInput,
    Global,
    rem,
} from "@mantine/core";
import { Dropzone, MIME_TYPES } from '@mantine/dropzone';
import { showNotification } from "@mantine/notifications";
import { useState, useEffect, Dispatch, SetStateAction } from "react";
import { CirclePlus, Edit, Tags, Trash, TableExport, TableImport, Settings, Photo, MessageCircle, Search, ArrowsLeftRight } from "tabler-icons-react";
import { containerStyles } from "../../../styles/container";
import { Item, Location_Item } from "../../../types/item";
import {ConfirmationModal} from "../../confirmation";

export const ItemRow = (
    {
        item,
        refresh
    }: {
        item: Item,
        refresh: () => Promise<void>
    }
) => {

    const doDelete = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/delete?id=${item.ID}`,
            {
                method: "DELETE",
                credentials: "include",
            }
        );

        if (response.ok) {
            showNotification({
                message: "Item deleted",
                color: "green",
                title: "Success",
            });
            await refresh();
            return;
        }

        showNotification({
            message: "Failed to delete item",
            color: "red",
            title: "Error",
        });
    };

    const editItem = async(name: string, sku: string, price: number, size: string, quantity: number) => {
        if(name == ""){
            name = item.name;
        }
        if(sku == ""){
            sku = item.sku;
        }
        if(price == -1){
            price = item.price;
        }
        if(size == ""){
            size = item.size;
        }
        if(quantity == -1){
            quantity = item.quantity;
        }
        const id = item.ID.toString();
        const stock_total = quantity;
        console.log("Run edit item id = %d, with name = %s, sku = %s, price = %d, size = %s, quantity = %d", item.ID, name, sku, price, size, quantity);
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/update`,
            {
                credentials: "include",
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    id,
                    name,
                    sku,
                    price,
                    size,
                    stock_total,
                }),
            }
        );
        if (response.ok) {
            showNotification({
                message: "Itemu updated",
                color: "green",
                title: "Success",
            });
            await refresh();
            return;
        }

        showNotification({
            message: "Failed to update the location",
            color: "red",
            title: "Error",
        });
    }

    const handleItemLocationDetailClick = async() => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/list-location-for-item?id=${item.ID}`,
            {
                method: "GET",
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: Location_Item[];
        } = await response.json();
        
        
        console.log(data.data);
        let locationItemMessage = "";
        let totalStored = 0;
        locationItemMessage = locationItemMessage + "List of locations which store: " + item.name + "\n";
        locationItemMessage += "----------\n";
        let idx = 0;
        for (idx=0; idx<data.data.length; idx++){
            locationItemMessage = locationItemMessage + (idx+1) + "." + data.data[idx].location_name + ": " +  data.data[idx].stock + "\n";
            totalStored += data.data[idx].stock;
        }
        const nonStored = item.quantity - totalStored;
        locationItemMessage += "-----\nTotal Stored: ";
        locationItemMessage += totalStored;
        locationItemMessage += "/";
        locationItemMessage += item.quantity;
        locationItemMessage += "\nNot Stored: "
        locationItemMessage += nonStored;
        locationItemMessage += "/";
        locationItemMessage += item.quantity;

        alert(locationItemMessage);
        
        return (
                <Table>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Quantity</th>
                    </tr>
                </thead>
                <tbody>
                    {data.data?.map((location_item) => (
                        <p>
                            <span>{location_item.location_name}</span>
                            <span>{location_item.stock}</span>
                        </p>
                    ))}
                </tbody>
            </Table>
        )

    }

    const [showConfirmationModal, setShowConfirmationModal] =
        useState<boolean>(false);
    const [editItemModal, setEditItemModal] =
        useState<boolean>(false);

    return (
        <>
            <tr>
                <td>{item.name}</td>
                <td>{item.sku || "None"}</td>
                <td>{item.price || "0"}</td>
                <td>
                    <Button onClick = {handleItemLocationDetailClick}>
                        {item.quantity || "0"}
                    </Button>
                </td>
                <td>{item.size || "None"}</td>
                <td>
                    <Group>
                    <HoverCard>
                        <HoverCard.Target>
                            <ActionIcon variant="default" onClick={() => setShowConfirmationModal(true)}>
                            <Trash />
                            </ActionIcon>
                        </HoverCard.Target>
                        <HoverCard.Dropdown>
                            <Text size="sm">
                                Delete this item
                            </Text>
                        </HoverCard.Dropdown>
                    </HoverCard>

                    <HoverCard>
                        <HoverCard.Target>
                        <ActionIcon variant="default" onClick={() => setEditItemModal(true)}>
                            <Edit />
                        </ActionIcon>
                        </HoverCard.Target>
                        <HoverCard.Dropdown>
                            <Text size="sm">
                                Edit this item
                            </Text>
                        </HoverCard.Dropdown>
                    </HoverCard>    
                        
                    </Group>
                </td>
            </tr>
            <EditItemModal opened={editItemModal} setOpened={setEditItemModal} command={editItem}/>
            <ConfirmationModal opened={showConfirmationModal} setOpened={setShowConfirmationModal} command={doDelete} message={"This action is not reversible. This will permanently delete the Item beyond recovery."}/>
        </>
    );
};
    
const  UploadCSVModal = (
{
    opened,
    setOpened,
    refresh,

}: {
    opened: boolean;
    setOpened: Dispatch<SetStateAction<boolean>>;
    refresh: () => Promise<void>;
}) => {
    const importCSV = async (file: File) => {
        const data = new FormData()
        await data.append("file", file)
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/import`,
            {
                credentials: "include",
                method: "POST",
                body: data
            }

        );
        await refresh();
        setOpened(false);
    };

    return (
        <>
            <Modal
                opened={opened}
                onClose={() => {
                    refresh();
                    setOpened(false);
                }}
            >
                <Dropzone
                    multiple={false}
                    onDrop={(file) => {importCSV(file[0])}}
                    maxSize={3 * 1024 ** 2}
                    accept={[MIME_TYPES.csv]}
                >
                    <Group position="center" spacing="xl" style={{ minHeight: 220, pointerEvents: 'none' }}>

                        <div>
                            <Text size="xl" inline>
                                Drag CSV here or click to select files
                            </Text>
                            <Text size="sm" color="dimmed" inline mt={7}>
                                Select CSV file containing you upload data
                            </Text>
                        </div>
                    </Group>
                </Dropzone>
            </Modal>
        </>
    );
}

export const EditItemModal = (
    {
        opened,
        setOpened,
        command,

    }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        command: (name: string, sku: string, price: number, size: string, quantity: number)=>void;
    }) => {

    const [name, setName] = useState('');
    const [sku, setSku] = useState('');
    const [price, setPrice] = useState(-1);
    const [size, setSize] = useState('');
    const [quantity, setQuantity] = useState(-1);


    return (
        <>
            <Modal
                title={"Edit Item"}
                opened={opened}
                onClose={() => {
                    setOpened(false);
                }}
            >
                <TextInput
                    required
                    label={"Name"}
                    placeholder="Left blank to use the curent name"
                    onChange={(e) => setSize(e.target.value)}
                />
                <Space h="md" />
                <TextInput
                    required
                    label={"SKU"}
                    placeholder="Left blank to use the curent SKU"
                    onChange={(e) =>
                        setQuantity(Number(e.target.value))
                    }
                />
                <TextInput
                    label="Price"
                    required
                    placeholder="Left blank to use the curent price"
                    type="number"
                    onChange={(e) =>
                        setPrice(Number(e.target.value))
                    }
                />
                <TextInput
                    required
                    label={"Size"}
                    placeholder="Left blank to use the curent size"
                    onChange={(e) => setSize(e.target.value)}
                />
                <Space h="md" />
                <TextInput
                    required
                    label={"Quantity"}
                    placeholder="Left blank to use the curent quantity"
                    type="number"
                    onChange={(e) =>
                        setQuantity(Number(e.target.value))
                    }
                />
                <Space h="md" />
                <Group position={"right"}>
                    <Button color="green" onClick={() => {command(name, sku, price, size, quantity); setOpened(false);}}>Confirm</Button>
                </Group>
            </Modal>
        </>
    );
}


const CreateItemModal = ({
    opened,
    setOpened,
    refresh,
}: {
    opened: boolean;
    setOpened: Dispatch<SetStateAction<boolean>>;
    refresh: () => Promise<void>;
}) => {
    const [name, setName] = useState("");
    const [sku, setSku] = useState("");
    const [price, setPrice] = useState(0);
    const [quantity, setQuantity] = useState(0);
    const [size, setSize] = useState("");


    const doCreate = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/add`,
            {
                credentials: "include",
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    name,
                    sku,
                    price,
                    quantity,
                    size,
                }),
            }
        );

        if (response.ok) {
            showNotification({
                color: "green",
                title: "Item created",
                message: "Item created successfully",
            });

            await refresh();
            setOpened(false);
            return;
        }

        const data = await response.json();

        showNotification({
            color: "red",
            title: "Error",
            message: data.message,
        });
    };

    return (
        <>
            <Modal
                opened={opened}
                onClose={() => {
                    refresh();
                    setOpened(false);
                }}
                title="Create Item"
            >
                <form
                    onSubmit={(e) => {
                        e.preventDefault();
                        doCreate();
                    }}
                >
                    <TextInput
                        label="Name"
                        required
                        placeholder="The name of the item"
                        onChange={(e) => setName(e.target.value)}
                    />
                    <Space h="md" />
                    <TextInput
                        label="SKU"
                        required
                        placeholder="XXXXXXXXX"
                        onChange={(e) => setSku(e.target.value)}
                    />
                    <Space h="md" />
                    <TextInput
                        label="Price"
                        required
                        placeholder="10"
                        type="number"
                        onChange={(e) =>
                            setPrice(Number(e.target.value))
                        }
                    />
                    <Space h="md" />
                    <TextInput
                        label="Quantity"
                        required
                        placeholder="10"
                        type="number"
                        onChange={(e) =>
                            setQuantity(Number(e.target.value))
                        }
                    />
                    <Space h="md" />
                    <TextInput
                        label="Size"
                        required
                        placeholder="10/XL"
                        onChange={(e) => setSize(e.target.value)}
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



export const ItemsManager = () => {
    const [items, setItems] = useState<Item[]>([]);
    const [loading, setLoading] = useState<boolean>(false);

    const [currentPage, setCurrentPage] = useState(1);
    const [totalPage, setTotalPage] = useState(1);

    const [nameQuery, setNameQuery] = useState("");
    const [nameQueryTyping, setNameQueryTyping] = useState("");

    const [skuQuery, setSkuQuery] = useState("");
    const [skuQueryTyping, setSkuQueryTyping] = useState("");

    

    const [showCreationModal, setShowCreationModal] = useState(false);
    const [showImportModal, setShowImportModal] = useState(false);



    const exportCSV = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/export?&name=${nameQuery}&sku=${skuQuery}`,
            {
                credentials: "include",
            }
        );

        // data is arraybuffer
        const data = await response.arrayBuffer();

        // convert to blob
        const blob = new Blob([data], { type: "text/csv" });

        // create a url from the blob
        const url = URL.createObjectURL(blob);

        // open the url with the pdf viewer
        window.open(url, "_blank");
    };


    const fetchItems = async () => {
        setLoading(true);

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/items/list?page=${currentPage}&name=${nameQuery}&sku=${skuQuery}`,
            {
                credentials: "include",
            }
        );

        setLoading(false);

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
            setTotalPage(data.data.totalPages);
        }
    };

    useEffect(() => {
        fetchItems();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        fetchItems();
    }, [nameQuery, skuQuery]);

    return (
        <>
            <CreateItemModal
                opened={showCreationModal}
                setOpened={setShowCreationModal}
                refresh={fetchItems}
            />
            <UploadCSVModal
                opened={showImportModal}
                setOpened={setShowImportModal}
                refresh={fetchItems}
            />
            <Box sx={containerStyles}>
                <Group position="apart">
                    <h3>Items</h3>
                    <Group spacing={0}>
                        <HoverCard>
                            <HoverCard.Target>
                                <ActionIcon
                                    sx={{
                                        height: "4rem",
                                        width: "4rem",
                                    }}
                                    onClick={() => setShowCreationModal(true)}
                                >
                                    <CirclePlus />
                                </ActionIcon>
                            </HoverCard.Target>
                            <HoverCard.Dropdown>
                                <Text size="sm">
                                    Add a new item
                                </Text>
                            </HoverCard.Dropdown>
                        </HoverCard>

                    </Group>
                </Group>
                <Space h="md" />
                <Group>
                    <TextInput
                        placeholder="Search Items"
                        onChange={(e: any) =>
                            setNameQueryTyping(e.target.value)
                        }
                    />
                    <TextInput
                        placeholder="Search SKU"
                        onChange={(e: any) =>
                            setSkuQueryTyping(e.target.value)
                        }
                    />
                    <Button
                        onClick={() => {
                            setNameQuery(nameQueryTyping);
                            setSkuQuery(skuQueryTyping);
                        }}
                        disabled={loading}
                    >
                        Search
                    </Button>
                </Group>


                <Space h="md" />
                <Table striped highlightOnHover>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>SKU</th>
                            <th>Price</th>
                            <th>Quantity</th>
                            <th>Size</th>
                            <th>Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {items.map((item) => (
                            <ItemRow key={item.ID} item={item} refresh={fetchItems} />
                        ))
                        }
                    </tbody>
                </Table>
                <Space h="md" />
                <Center>
                    {loading ? (
                        <Loader variant="dots" />
                    ) : (
                        <Pagination
                            boundaries={3}
                            withEdges
                            value={currentPage}
                            total={totalPage}
                            onChange={setCurrentPage}
                        />
                    )}
                </Center>
                <Space h="md" />
                <Group position="apart">
                    <Group spacing={0}>
                        <ActionIcon
                            sx={{
                                height: "2.5rem",
                                width: "2.5rem",
                            }}
                            onClick={exportCSV}
                        >
                            <TableExport size={"1.5rem"}/>
                        </ActionIcon>
                        <ActionIcon
                            sx={{
                                height: "2.5rem",
                                width: "2.5rem",
                            }}
                            onClick={() => {setShowImportModal(true)}}
                        >
                            <TableImport size={"1.5rem"}/>
                        </ActionIcon>
                    </Group>
                    <Button
                        onClick={fetchItems}
                        disabled={loading}
                    >
                        Refresh
                    </Button>
                </Group>
            </Box>
        </>
    );
};
function useStyles(): { classes: any; cx: any; } {
    throw new Error("Function not implemented.");
}

