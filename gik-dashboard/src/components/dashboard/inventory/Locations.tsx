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
import {Client} from "../../../types/client";
import {ConfirmationModal} from "../../confirmation";

import type { Location } from "../../../types/location";

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

export const LocationManager = () => {
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
            `${process.env.REACT_APP_API_URL}/location/list?page=${currentPage}&name=${nameQuery}&sku=${skuQuery}`,
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
            //console.log(data.data.data)
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
                            setDescriptionQuery(e.target.value)
                        }
                    />
                    {/* <MultiSelect
                        data={tags}
                        placeholder="Search Tags"
                        clearButtonProps={{ 'aria-label': 'Clear selection' }}
                        clearable
                        searchable
                        onChange={(e: any) => {
                                setTagsQueryTyping(e)
                            }
                        }
                    /> */}
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
                        </tr>
                    </thead>
                    <tbody>
                        {locations.map((locaton) => (
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
