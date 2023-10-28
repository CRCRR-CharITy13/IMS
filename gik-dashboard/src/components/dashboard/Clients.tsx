import {
    Box,
    Table,
    Space,
    Group,
    Button,
    TextInput,
    Modal,
    ActionIcon, 
    Text
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { showNotification } from "@mantine/notifications";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { CirclePlus, TableExport, TableImport, Trash} from "tabler-icons-react";
import { containerStyles } from "../../styles/container";
import { Client } from "../../types/client";
import { Dropzone, MIME_TYPES } from "@mantine/dropzone";
import { ConfirmationModal } from "../confirmation";
import { Checkbox } from '@mantine/core';

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
            `${process.env.REACT_APP_API_URL}/client/import`,
            {
                credentials: "include",
                method: "POST",
                body: data
            }

        );
        await refresh();
        setOpened(false)
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
// Select Client Component using checkbox

// const SelectClientComponent = ({
//     client,
//     refresh,
// }: {
//     client: Client;
//     refresh: () => Promise<void>;
// }) => {
//     const SelectClient = async ()=> {
//         const response = await fetch(
//         `${process.env.REACT_APP_API_URL}/client/select?id=${client.ID}`,
//         {
//             method: "SELECT",
//             credentials: "include",
//         }
//     );
//     }
// }


const ClientComponent = ({
    client,
    refresh,
}: {
    client: Client;
    refresh: () => Promise<void>;
}) => {
    const doDelete= async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/client/delete?id=${client.ID}`,
            {
                method: "DELETE",
                credentials: "include",
            }
        );

        if (response.ok) {
            showNotification({
                message: "Client deleted",
                color: "green",
                title: "Success",
            });
            await refresh();
            return;
        }

        showNotification({
            message: "Failed to delete client",
            color: "red",
            title: "Error",
        });
    };

    const [showConfirmationModal, setShowConfirmationModal] =
        useState<boolean>(false);

    return (
        <>
            <tr>
                <td>{client.name}</td>
                <td>{client.contact}</td>
                <td>{client.phone}</td>
                <td>{client.email}</td>
                <td>{client.address}</td>
                <td>{client.balance}</td>
                <td>
                    <ActionIcon variant="default" onClick={() => setShowConfirmationModal(true)}>
                        <Trash />
                    </ActionIcon>
                </td>
                <td>
                    <Checkbox defaultValue="disabled"
                        // labelPosition="left"
                        color="teal"
                        size="md"
                        onClick={() => showConfirmationModal}
                    >
                    </Checkbox>
                </td>                
            </tr>
            <ConfirmationModal opened={showConfirmationModal} setOpened={setShowConfirmationModal} command={doDelete} message={"This action is not reversible. This will permanently delete the client beyond recovery."}/>
        </>
    );
};

const CreateClientModal = ({
    opened,
    setOpened,
    refresh,
}: {
    opened: boolean;
    setOpened: Dispatch<SetStateAction<boolean>>;
    refresh: () => Promise<void>;
}) => {
    const form = useForm({
        initialValues: {
            name: "",
            contact: "",
            phone: "",
            email: "",
            address: "",
            balance: 0,
        },
    });

    const doCreate = async () => {
        form.values.balance = Number(form.values.balance);

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/client/add`,
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify(form.values),
            }
        );

        const data: {
            success: boolean;
            message: string;
        } = await response.json();

        if (data.success) {
            showNotification({
                color: "green",
                title: "Client added",
                message: data.message,
            });

            await refresh();
            setOpened(false);
            return;
        }

        showNotification({
            color: "red",
            title: "Unable to add client",
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
                title="Create Client"
            >
                <form
                    onSubmit={(e) => {
                        e.preventDefault();
                        doCreate();
                    }}
                >
                    <TextInput
                        label="Organization name"
                        required
                        placeholder="Gifts In Kind"
                        {...form.getInputProps("name")}
                    />
                    <Space h="md" />
                    <TextInput
                        label="Contact Person"
                        required
                        placeholder="Mr. John Doe"
                        {...form.getInputProps("contact")}
                    />
                    <Space h="md" />
                    <Group grow>
                        <TextInput
                            label="Phone"
                            required
                            type="tel"
                            placeholder="+1 (555) 555-5555"
                            {...form.getInputProps("phone")}
                        />
                        <TextInput
                            label="Email"
                            required
                            placeholder="john@giftsinkindottawa.ca"
                            type="email"
                            {...form.getInputProps("email")}
                        />
                    </Group>
                    <Space h="md" />
                    <TextInput
                        label="Address"
                        required
                        placeholder="123 Main St"
                        {...form.getInputProps("address")}
                    />
                    <Space h="md" />
                    <TextInput
                        label="Balance"
                        required
                        placeholder="50"
                        type="number"
                        {...form.getInputProps("balance")}
                    />
                    <Space h="md" />
                    <Group position="right">
                        <Button type="submit">
                            Create
                        </Button>
                    </Group>
                </form>
            </Modal>
        </>
    );
};

export const ClientManager = () => {
    const [clients, setClients] = useState<Client[]>([]);

    const [loading, setLoading] = useState<boolean>(true);

    const [nameQuery, setNameQuery] = useState<string>("");

    const [contactQuery, setContactQuery] = useState<string>("");

    const [phoneQuery, setPhoneQuery] = useState<string>("");

    const [emailQuery, setEmailQuery] = useState<string>("");

    const [addressQuery, setAddressQuery] = useState<string>("");

    const [showClientCreationModal, setShowClientCreationModal] =
        useState<boolean>(false);

    const [showImportModal, setShowImportModal] =
        useState<boolean>(false);

    const exportCSV = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/client/export?name=${nameQuery}&contact=${contactQuery}&phone=${phoneQuery}&email=${emailQuery}&address=${addressQuery}`,
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

    const fetchClients = async () => {
        setLoading(true);

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/client/list?name=${nameQuery}&contact=${contactQuery}&phone=${phoneQuery}&email=${emailQuery}&address=${addressQuery}`,
            {
                credentials: "include",
            }
        );

        setLoading(false);

        const data: {
            success: boolean;
            data: Client[];
        } = await response.json();

        if (data.success) {
            setClients(data.data);
        }
        console.log(data.data);
    };

    useEffect(() => {
        fetchClients();
    }, []);

    return (
        <>
            <Box sx={containerStyles}>
                <Group position="apart">
                    <h3>Clients</h3>
                    <ActionIcon
                        sx={{
                            height: "4rem",
                            width: "4rem",
                        }}
                        onClick={() => setShowClientCreationModal(true)}
                    >
                        <CirclePlus />
                    </ActionIcon>
                </Group>
                <Space h="md" />
                <Group position="apart">
                    <TextInput
                        placeholder="Organization name"
                        onChange={(e) => setNameQuery(e.target.value)}
                    />
                    <TextInput
                        placeholder="Contact"
                        onChange={(e) => setContactQuery(e.target.value)}
                    />
                    <TextInput
                        placeholder="Phone"
                        onChange={(e) => setPhoneQuery(e.target.value)}
                    />
                    <TextInput
                        placeholder="Email"
                        onChange={(e) => setEmailQuery(e.target.value)}
                    />
                    <TextInput
                        placeholder="Address"
                        onChange={(e) => setAddressQuery(e.target.value)}
                    />   
                </Group>
                <Space h="md"/>
                <Group position="apart">
                    <Button
                        disabled={loading}
                        onClick={() => fetchClients()}
                    >
                        Filter
                    </Button>
                    {/* <Button color="red" onClick={() => ConfirmationModal}>
                    Delete Clients
                    </Button> */}
                    </Group>

                <Space h="md" />
                <Table>
                    <thead>
                        <tr>
                            <th>Organization Name</th>
                            <th>Contact</th>
                            <th>Phone</th>
                            <th>Email</th>
                            <th>Address</th>
                            <th>Balance</th>
                            <th>Actions</th>
                            <th>
                                {/* <Checkbox size="md" 
                                    onClick={() => getSelection()} >
                             
                                </Checkbox> */}
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        {clients.map((client) => (
                            <ClientComponent
                                refresh={fetchClients}
                                key={client.ID}
                                client={client}
                            />
                        ))}
                    </tbody>
                </Table>
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
                </Group>
            </Box>
            <CreateClientModal
                opened={showClientCreationModal}
                setOpened={setShowClientCreationModal}
                refresh={fetchClients}
            />
            <UploadCSVModal
                opened={showImportModal}
                setOpened={setShowImportModal}
                refresh={fetchClients}
            />
        </>
    );
};

export default ClientManager;