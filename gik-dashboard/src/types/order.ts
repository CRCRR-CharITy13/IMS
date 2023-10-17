export interface Order {
    ID: number;
    timestamp: string;
    clientName: string;
    signedBy: string;
    totalQuantity: number;
}

export interface OrderItem {
    ID: number;
    name: string;
    sku: string;
    size: string;
    price: number;
    quantity: number;
    totalValue: number;
}