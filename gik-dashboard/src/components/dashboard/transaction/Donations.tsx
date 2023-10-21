import { 
    Table 
} from "@mantine/core";
export const DonationManager = () => {
    return (
        <>
        <h3>Donations</h3>
        <Table>
            <thead>
                <tr>
                    <th>
                        <th> Time </th>
                        <th> Signed By </th>
                        <th> Donator </th>
                        <th> Total Value </th>
                        <th> Action </th>
                    </th>
                </tr>
            </thead>
        </Table>
        </>
    )
}