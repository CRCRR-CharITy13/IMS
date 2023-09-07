import {
    Table
} from "@mantine/core";

import { Item } from "../../../types/item";
const exampleData = [
    {
        id: 1,
        name: "men T-Shirt",
        sku: "B15TSH",
        size: "L",
        quantity: 100,
        price: 40
    },
    {
        id: 1,
        name: "women T-Shirt",
        sku: "B19TSH",
        size: "M",
        quantity: 10,
        price: 50
    }
];

const SimpleItemRow = ({ item } : {item: Item}) => {
    return (
        <>
            <tr>
                <td>{item.id}</td>
                <td>{item.name}</td>
                <td>{item.sku}</td>
                <td>{item.size}</td>
                <td>{item.quantity}</td>
                <td>{item.price}</td>
            </tr>
        </>
    )
}
export const ItemsManager = () => {
return (
<>
<Table>
    <thead>
        <tr>
            <th>ID</th>
            <th>Name</th>
            <th>SKU</th>
            <th>Size</th>
            <th>Quantity</th>
            <th>Price</th>
        </tr>
    </thead>
    <tbody>
        {exampleData.map((item) => (
            <SimpleItemRow key = {item.id } item = {item} />
        ))}
    </tbody>
</Table>
</>
);
};