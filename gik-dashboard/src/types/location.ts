import { Item } from "./item";

export interface Location {
    ID: number;
    id: number;
    name: string;
    letter: string;
    itemId: number;
    product: Item;
    sku: string;
    productName: string;
}