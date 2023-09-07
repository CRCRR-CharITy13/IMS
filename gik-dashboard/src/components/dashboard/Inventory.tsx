// import { Divider, Space } from "@mantine/core";
import { ItemsManager } from "./inventory/Items";
// import { LocationsManager } from "./inventory/Locations";

const Inventory = () => {
    return (
        <>
            <ItemsManager/>
        </>
    );
};

export default Inventory;


// const Inventory = () => {
//     return (
//         <>
//             <ItemsManager/>
//             <Space h="md" />
//             <Space h="md" />
//             <Divider />
//             <LocationsManager/>
//         </>
//     );
// };