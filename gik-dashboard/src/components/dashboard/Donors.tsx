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
import { Donor } from "../../types/donor";
import { Dispatch, SetStateAction, useState, useEffect } from "react";
import { CirclePlus, TableExport, TableImport, Trash, Edit} from "tabler-icons-react";
import { ConfirmationModal } from "../confirmation";
import { containerStyles } from "../../styles/container";

export const UpdateDonorModal = (
    {
        opened,
        setOpened,
        command,

    }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        command: (name : string, phone : string, email : string, address : string)=>void;
    }) => {

    const [name, setName] = useState('');
    const [phone, setPhone] = useState('');
    const [email, setEmail] = useState('');
    const [address, setAddress] = useState('');
    
    return (
        <>
            <Modal
                title={"Update Donor"}
                opened={opened}
                onClose={() => {
                    setOpened(false);
                }}
            >
                <TextInput
                    required
                    label={"Name"}
                    placeholder="Left blank to use the current name"
                    onChange={(e) => setName(e.target.value)}
                />
                <Space h="md" />
                <TextInput
                    required
                    label={"Phone"}
                    placeholder="Left blank to use the current description"
                    onChange={(e) => setPhone(e.target.value)}
                />
                <TextInput
                    required
                    label={"Email"}
                    placeholder="Left blank to use the current name"
                    onChange={(e) => setEmail(e.target.value)}
                />
                <Space h="md" />
                <TextInput
                    required
                    label={"Address"}
                    placeholder="Left blank to use the current description"
                    onChange={(e) => setAddress(e.target.value)}
                />

                <Space h="md" />
                <Group position={"right"}>
                    <Button color="green" onClick={() => {command(name, phone, email, address); setOpened(false);}}>Confirm</Button>
                </Group>
            </Modal>
        </>
    );
}

const DonorComponent = ({
    donor,
    refresh,
} : {
    donor : Donor;
    refresh: () => Promise<void>;
}) => {

    const [showConfirmationModal, setShowConfirmationModal] =
        useState<boolean>(false);
    
    const doDelete= async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donor/delete?id=${donor.ID}`,
            {
                method: "DELETE",
                credentials: "include",
            }
        );

        if (response.ok) {
            showNotification({
                message: "Donor deleted",
                color: "green",
                title: "Success",
            });
            await refresh();
            return;
        }

        showNotification({
            message: "Failed to delete donor",
            color: "red",
            title: "Error",
        });
    };

    const doUpdate = async(name: string, phone : string, email : string, address : string) => {
        if(name == ""){
            name = donor.name;
        }
        if(phone == ""){
            phone = donor.phone;
        }
        if(email == ""){
            email = donor.email;
        }
        if(address == ""){
            address = donor.address;
        }
        const id = donor.ID.toString();
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donor/update?id=${donor.ID}`,
            {
                credentials: "include",
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    name,
                    phone,
                    email,
                    address,
                }),
            }
        );
        if (response.ok) {
            showNotification({
                message: "Donor updated",
                color: "green",
                title: "Success",
            });
            await refresh();
            return;
        }

        showNotification({
            message: "Failed to update the donor",
            color: "red",
            title: "Error",
        });
    }

    const [updateDonorModal, setUpdateDonorModal] = useState<boolean>(false);

    return (
    <>
        <tr>
            <td>{donor.name}</td>
            <td>{donor.phone}</td>
            <td>{donor.email}</td>
            <td>{donor.address}</td>
            <td>
                <Group>
                <ActionIcon variant="default" onClick={() => setShowConfirmationModal(true)}>
                    <Trash />
                </ActionIcon>
                <ActionIcon variant="default" onClick={() => setUpdateDonorModal(true)}>
                    <Edit />
                </ActionIcon>
                </Group>
            </td>         
        </tr>
        <UpdateDonorModal opened={updateDonorModal} setOpened={setUpdateDonorModal} command={doUpdate}/>
        <ConfirmationModal opened={showConfirmationModal} setOpened={setShowConfirmationModal} command={doDelete} message={"This action is not reversible. This will permanently delete the client beyond recovery."}/>
    </>
    );
}


const CreateDonorModal = ({
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
            phone: "",
            email: "",
            address: "",
        },
    });

    const doCreate = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donor/add`,
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
                title: "Donor added",
                message: data.message,
            });

            await refresh();
            setOpened(false);
            return;
        }

        showNotification({
            color: "red",
            title: "Unable to add donor",
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
                title="Create Donor"
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
                        placeholder="Gifts In Kind"
                        {...form.getInputProps("name")}
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

export const DonorManager = () => {
    const [donors, setDonors] = useState<Donor[]>([]) 
    const [loading, setLoading] = useState<boolean>(true);
    const [nameQuery, setNameQuery] = useState<string>("");
    const [phoneQuery, setPhoneQuery] = useState<string>("");
    const [emailQuery, setEmailQuery] = useState<string>("");
    const [addressQuery, setAddressQuery] = useState<string>("");
    const [showDonorCreationModal, setShowDonorCreationModal] =
    useState<boolean>(false);


    // to fetch donor
    const fetchDonors = async () => {
        setLoading(true);

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/donor/list?name=${nameQuery}&address=${addressQuery}&phone=${phoneQuery}&email=${emailQuery}&address=${addressQuery}`,
            {
                credentials: "include",
            }
        );

        setLoading(false);

        const data: {
            success: boolean;
            data: Donor[];
        } = await response.json();

        if (data.success) {
            setDonors(data.data);
        }
        console.log(data.data);
    };

    useEffect(() => {
        fetchDonors();
    }, []);

    return (
        <>
            <Box sx={containerStyles}>
                    <Group position="apart">
                        <h3>Donors</h3>
                        <ActionIcon
                            sx={{
                                height: "4rem",
                                width: "4rem",
                            }}
                            onClick={() => setShowDonorCreationModal(true)}
                        >
                            <CirclePlus />
                        </ActionIcon>
                    </Group>
                    <Space h="md" />
                    <Group position="apart">
                        <TextInput
                            placeholder="Name"
                            onChange={(e) => setNameQuery(e.target.value)}
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
                            onClick={() => fetchDonors()}
                        >
                            Filter
                        </Button>
                        </Group>

                    <Space h="md" />
                <Table>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Phone</th>
                            <th>Email</th>
                            <th>Address</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {donors.map((donor) => (
                            <DonorComponent 
                                refresh = {fetchDonors}
                                key = {donor.ID}
                                donor = {donor}
                            />
                        ))}
                    </tbody>
                </Table>
            </Box>
            <CreateDonorModal
                opened={showDonorCreationModal}
                setOpened={setShowDonorCreationModal}
                refresh={fetchDonors}
            />
        </>
    )
}

export default DonorManager;