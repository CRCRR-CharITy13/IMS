import { Divider, Space } from "@mantine/core";
import { OrderManager } from "./transaction/Orders"
import { DonationManager } from "./transaction/Donations";


const Transactions = () => {
    return (
        <>
            <OrderManager/>
            <Space h="md"/>
            <Divider/>
            <DonationManager/>
        </>
    );
};
export default Transactions;