export interface Item {
    ID: number;
    name: string;
    sku: string;
    price: number;
    quantity: number;
    size: string;
}

export interface Location_Item {
    location_name : string;
    stock : number;
}