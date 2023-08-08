export interface Transaction {
    ID: number;
    timestamp: number;
    clientId: number;
    type: boolean;
    signerId: number;
    totalQuantity: number;
}

export interface TransactionItem {
    ID: number;
    name: string;
    sku: string;
    size: string;
    price: number;
    quantity: number;
    totalValue: number;
}