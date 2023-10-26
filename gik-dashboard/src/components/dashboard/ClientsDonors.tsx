import { ClientManager } from "./Clients";
import { DonorManager } from "./Donors"
import { Divider, Space } from "@mantine/core";

const ClientsDonors = () => {
    return (
        <>
            <ClientManager />
            <Space h="md"/>
            <Space h="md"/>
            <Divider/>
            <DonorManager />
        </>
    );
};

export default ClientsDonors;