export interface Order {
    ID: number;
    createdTime: string;
    clientId: number;
    clientName: string;
    signerId : number;
    signedBy: string;
    totalCost: number;
}

export interface OrderItem {
    ID: number;
    name: string;
    sku: string;
    size: string;
    price: number;
    quantity: number;
    totalCost: number;
}

export interface AddOrderResponse {
    ID: number;
    itemSKUName : string;
    locationName : string;
    quantity : number;
}