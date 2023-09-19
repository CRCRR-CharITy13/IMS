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
import { CirclePlus, ClipboardPlus, Edit, Tags, Trash, TableExport, TableImport, Settings, Photo, MessageCircle, Search, ArrowsLeftRight } from "tabler-icons-react";
import { containerStyles } from "../../../styles/container";
import { Item } from "../../../types/item";
import { ConfirmationModal } from "../../confirmation";

import { Location } from "../../../types/location";

import { Item_Location } from "../../../types/location";

export const EditLocationModal = (
    {
        opened,
        setOpened,
        command,

    }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        command: (name : string, description : string)=>void;
    }) => {

    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    
    return (
        <>
            <Modal
                title={"Edit Location"}
                opened={opened}
                onClose={() => {
                    setOpened(false);
                }}
            >
                <TextInput
                    required
                    label={"Name"}
                    placeholder="Left blank to use the current value"
                    onChange={(e) => setName(e.target.value)}
                />
                <Space h="md" />
                <TextInput
                    required
                    label={"Description"}
                    placeholder="Left blank to use the current value"
                    onChange={(e) =>
                        setDescription(e.target.value)}
                />

                <Space h="md" />
                <Group position={"right"}>
                    <Button color="green" onClick={() => {command(name, description); setOpened(false);}}>Confirm</Button>
                </Group>
            </Modal>
        </>
    );
}

export const AddItemToLocationModal = (
    {
        opened,
        setOpened,
        command,

    }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        command: (size: string, quantity: number)=>void;
    }) => {

    const [itemSKU, setItemSKU] = useState('');
    const [quantity, setQuantity] = useState(0);


    return (
        <>
            <Modal
                title={"Add Item to Location"}
                opened={opened}
                onClose={() => {
                    setOpened(false);
                }}
            >
                <TextInput
                    required
                    label={"Item SKU"}
                    placeholder="XXXXXXXXX"
                    onChange={(e) => setItemSKU(e.target.value)}
                />
                <Space h="md" />
                <TextInput
                    required
                    label={"Quantity"}
                    placeholder="10"
                    type="number"
                    onChange={(e) =>
                        setQuantity(Number(e.target.value))
                    }
                />
                <Space h="md" />
                <Group position={"right"}>
                    <Button color="green" onClick={() => {command(itemSKU, quantity); setOpened(false);}}>Confirm</Button>
                </Group>
            </Modal>
        </>
    );
}

export const LocationRow = (
    {
        location,
        refresh
    }: {
        location: Location,
        refresh: () => Promise<void>
    }
) => {
    const addItemToLocation = async (itemSKU: string, quantity: number) => {
        const locationID = location.ID
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/location/add-item-to-location`,
            {
                credentials: "include",
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    itemSKU,
                    locationID,
                    quantity,
                }),
            }
        );
        if (response.ok) {
            showNotification({
                message: "item added to location",
                color: "green",
                title: "Success",
            });
            await refresh();
            return;
        }
        console.log("add %d items with SKU = %s to box id %d", quantity, itemSKU, location.ID)
        
        
        showNotification({
            message: "Failed to add size",
            color: "red",
            title: "Error",
        });
    };
    const handleLocationDetailClick = async() => {

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/location/list-item-in-location?id=${location.ID}`,
            {
                method: "GET",
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: Item_Location[];
        } = await response.json();
        
        
        console.log(data.data);
        let itemLocationMessage = "";
        itemLocationMessage = itemLocationMessage + "List of items in " + location.name + "\n";
        itemLocationMessage += "----------\n";
        let idx = 0;
        for (idx=0; idx<data.data.length; idx++){
            itemLocationMessage = itemLocationMessage + (idx+1) + "." + data.data[idx].item_name + ": " +  data.data[idx].stock + "\n";
        }
        alert(itemLocationMessage);
        return (
                <Table>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Quantity</th>
                    </tr>
                </thead>
                <tbody>
                    {data.data?.map((item_location) => (
                        <p>
                            <span>{item_location.item_name}</span>
                            <span>{item_location.stock}</span>
                        </p>
                    ))}
                </tbody>
            </Table>
        )

    }
    const doDelete = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/location/delete?id=${location.ID}`,
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

    const editLocation = async(name: string, description : string) => {
        if(name == ""){
            name = location.name;
        }
        if(description == ""){
            description = location.description;
        }
        const id = location.ID.toString();
        console.log("Run edit location id = %d, with name = %s, description = %s", location.ID, name, description);
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/location/update`,
            {
                credentials: "include",
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    id,
                    name,
                    description,
                }),
            }
        );
        if (response.ok) {
            showNotification({
                message: "Location changed",
                color: "green",
                title: "Success",
            });
            await refresh();
            return;
        }

        showNotification({
            message: "Failed to change the location",
            color: "red",
            title: "Error",
        });
    }

    const [showConfirmationModal, setShowConfirmationModal] = useState<boolean>(false);
    const [editLocationModal, setEditLocationModal] = useState<boolean>(false);
    const [addItemToLocationModal, setAddItemToLocationModal] = useState<boolean>(false);
    return (
        <>
            <tr>
                <td>{location.name || "None"}</td>
                <td>{location.description || "None"}</td>
                <td>
                    <Button onClick = {handleLocationDetailClick}>
                        {location.total_item || 0}
                    </Button>
                </td>
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
                                Delete this location
                            </Text>
                        </HoverCard.Dropdown>
                    </HoverCard>

                    <HoverCard>
                        <HoverCard.Target>
                        <ActionIcon variant = "default" onClick={() => setEditLocationModal(true)}>
                            <Edit />
                        </ActionIcon>
                        </HoverCard.Target>
                        <HoverCard.Dropdown>
                            <Text size="sm">
                                Edit this location
                            </Text>
                        </HoverCard.Dropdown>
                    </HoverCard>

                    <HoverCard>
                        <HoverCard.Target>
                        <ActionIcon variant="default" onClick={() => setAddItemToLocationModal(true)}>
                            <ClipboardPlus />
                            </ActionIcon>
                        </HoverCard.Target>
                        <HoverCard.Dropdown>
                            <Text size="sm">
                                Add item to this location
                            </Text>
                        </HoverCard.Dropdown>
                    </HoverCard>

                    </Group>
                </td>
            </tr>
            <EditLocationModal opened={editLocationModal} setOpened={setEditLocationModal} command={editLocation}/>
            <ConfirmationModal opened={showConfirmationModal} setOpened={setShowConfirmationModal} command={doDelete} message={"This action is not reversible. This will permanently delete the Location beyond recovery."}/>
            <AddItemToLocationModal opened={addItemToLocationModal} setOpened={setAddItemToLocationModal} command={addItemToLocation}/>
        </>
    );
};

const CreateLocationModal = ({
    opened,
    setOpened,
    refresh,
}: {
    opened: boolean;
    setOpened: Dispatch<SetStateAction<boolean>>;
    refresh: () => Promise<void>;
}) => {
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const doCreate = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/location/add`,
            {
                credentials: "include",
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    name,
                    description,
                }),
            }
        );

        if (response.ok) {
            showNotification({
                color: "green",
                title: "Location created",
                message: "Location created successfully",
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
                title="Create Location"
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
                        placeholder="The name of the location"
                        onChange={(e) => setName(e.target.value)}
                    />
                    <Space h="md" />
                    <TextInput
                        label="Description"
                        required
                        placeholder="The description of the location"
                        onChange={(e) => setDescription(e.target.value)}
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

export const LocationsManager = () => {
    const [locations, setLocations] = useState<Location[]>([]);
    const [loading, setLoading] = useState<boolean>(false);

    const [currentPage, setCurrentPage] = useState(1);
    const [totalPage, setTotalPage] = useState(1);

    const [nameQuery, setNameQuery] = useState("");
    const [nameQueryTyping, setNameQueryTyping] = useState("");

    const [descriptionQuery, setDescriptionQuery] = useState("");
    const [descriptionQueryTyping, setDescriptionQueryTyping] = useState("");

    const [showCreationModal, setShowCreationModal] = useState(false);

    

    const fetchLocations = async () => {
        setLoading(true);

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/location/list?page=${currentPage}&name=${nameQuery}&description=${descriptionQuery}`,
            {
                credentials: "include",
            }
        );

        setLoading(false);

        const data: {
            success: boolean;
            message: string;
            data: {
                data: Location[];
                total: number;
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            console.log("Success loading item data");
            setLocations(data.data.data);
            console.log(data.data.data);
            setTotalPage(data.data.totalPages);
        }
    };

    useEffect(() => {
        fetchLocations();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        fetchLocations();
    }, [nameQuery, descriptionQuery]);

    return (
        <>
            <CreateLocationModal
                opened={showCreationModal}
                setOpened={setShowCreationModal}
                refresh={fetchLocations}
            />
            <Box sx={containerStyles}>
                <Group>
                <Group position="apart">
                    <h3>Locations</h3>
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
                                    Add a new location
                                </Text>
                            </HoverCard.Dropdown>
                        </HoverCard>
                    </Group>
                </Group>
                <Space h="md" />
                <Group>
                    <TextInput
                        placeholder="Search Name"
                        onChange={(e: any) =>
                            setNameQueryTyping(e.target.value)
                        }
                    />
                    <TextInput
                        placeholder="Search Description"
                        onChange={(e: any) =>
                            setDescriptionQueryTyping(e.target.value)
                        }
                    />
                    <Button
                        onClick={() => {
                            setNameQuery(nameQueryTyping);
                            setDescriptionQuery(descriptionQueryTyping);
                        }}
                        disabled={loading}
                    >
                    Search
                    </Button>
                </Group>
            </Group>
            <Space h="md" />
            <Table striped highlightOnHover>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Description</th>
                        <th>Nb of Items</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {locations.map((location) => (
                        <LocationRow key={location.ID} location={location} refresh={fetchLocations} />
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
            
                <Button
                    onClick={fetchLocations}
                    disabled={loading}
                >
                    Refresh
                </Button>
            </Box>
        </>
    );
};
function useStyles(): { classes: any; cx: any; } {
    throw new Error("Function not implemented.");
}

export default LocationsManager;
