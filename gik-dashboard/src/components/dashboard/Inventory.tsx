import { ItemsManager } from "./inventory/Items";
import { LocationsManager } from "./inventory/Locations";

const Inventory = () => {
    return (
        <>
            <ItemsManager/>
            <LocationsManager/>
        </>
    );
};

export default Inventory;
