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
import { Item } from "../../../types/item";
import { ConfirmationModal } from "../../confirmation";

import { Location } from "../../../types/location";

export const EditLocationModal = (
    {
        opened,
        setOpened,
        // command,

    }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        //command: (location : Location)=>void;
    }) => {

    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [locationInstance, setLocationInstance] = useState();
    

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
                    placeholder="The name of the location"
                    onChange={(e) => setName(e.target.value)}
                />
                <Space h="md" />
                <TextInput
                    required
                    label={"Description"}
                    placeholder="The description of the location"
                    onChange={(e) =>
                        setDescription(e.target.value)}
                />

                <Space h="md" />
                {/* <Group position={"right"}>
                    <Button color="green" onClick={() => {command(location); setOpened(false);}}>Confirm</Button>
                </Group> */}
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

    const editLocation = async(location : Location) => {
        console.log("Run edit location id = %d", location.ID);
    }

    const [editLocationModal, setEditLocationModal] = useState<boolean>(false);
    return (
        <>
            <tr>
                <td>{location.name || "None"}</td>
                <td>{location.description || "None"}</td>
                <td>
                    <Group>
                    <HoverCard>
                        <HoverCard.Target>
                            <ActionIcon>
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
                    </Group>
                </td>
            </tr>
            <EditLocationModal opened={editLocationModal} setOpened={setEditLocationModal} />
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
            <Box sx={containerStyles}>
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
                <Space h="md" />
                <Table striped highlightOnHover>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Description</th>
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
