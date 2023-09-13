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

export const LocationRow = (
    {
        location,
        refresh
    }: {
        location: Location,
        refresh: () => Promise<void>
    }
) => {

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
                                Delete this item
                            </Text>
                        </HoverCard.Dropdown>
                    </HoverCard>

                    <HoverCard>
                        <HoverCard.Target>
                        <ActionIcon>
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
