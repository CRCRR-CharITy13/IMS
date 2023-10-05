export interface Location {
    ID: number;
    name: string;
    description: string;
    total_item : number;
}

export interface Item_Location {
    item_sku    : string;
    item_name : string;
    stock : number;
}