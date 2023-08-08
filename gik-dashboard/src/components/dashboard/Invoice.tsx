import {
    ActionIcon,
    Box,
    Button,
    Divider,
    Group,
    InputWrapper,
    Modal,
    Space,
    Table,
    Textarea,
    TextInput,
} from "@mantine/core";
import { Dispatch, SetStateAction, useState } from "react";
import { Trash } from "tabler-icons-react";
import { containerStyles } from "../../styles/container";

interface invoiceItem {
    name: string;
    size: string;
    sku: string;
    price: number;
    quantity: number;
}

interface invoiceData {
    name: string;
    address: string;
    data: invoiceItem[];
}

const InvoiceItemRow = ({
    item,
    invoiceItems,
    setInvoiceItems,
}: {
    item: invoiceItem;
    invoiceItems: invoiceItem[];
    setInvoiceItems: Dispatch<SetStateAction<invoiceItem[]>>;
}) => {
    return (
        <>
            <tr>
                <td>{item.name}</td>
                <td>{item.size}</td>
                <td>{item.sku}</td>
                <td>{item.price}</td>
                <td>{item.quantity}</td>
                <td>
                    <ActionIcon
                        onClick={() => {
                            // remove item from invoice items
                            setInvoiceItems(
                                invoiceItems.filter((i) => i.sku !== item.sku)
                            );
                        }}
                        variant="default"
                    >
                        <Trash />
                    </ActionIcon>
                </td>
            </tr>
        </>
    );
};

const InvoiceConfiguratorModal = ({
    opened,
    setOpened,
}: {
    opened: boolean;
    setOpened: Dispatch<SetStateAction<boolean>>;
}) => {
    const [customerName, setCustomerName] = useState("");
    const [customerAddress, setCustomerAddress] = useState("");

    const [invoiceItems, setInvoiceItems] = useState([] as invoiceItem[]);

    const [itemName, setItemName] = useState("");
    const [itemDesc, setItemDesc] = useState("");
    const [itemSku, setItemSku] = useState("");
    const [itemPrice, setItemPrice] = useState(0);
    const [itemQuantity, setItemQuantity] = useState(0);

    const generateInvoice = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/invoice/generate`,
            {
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                method: "POST",
                body: JSON.stringify({
                    name: customerName,
                    address: customerAddress,
                    data: invoiceItems,
                }),
            }
        );

        const data = await response.arrayBuffer();

        const blob = new Blob([data], { type: "application/pdf" });

        const url = URL.createObjectURL(blob);

        window.open(url);
    };

    return (
        <>
            <Modal
                size="xl"
                title="Configure Invoice"
                opened={opened}
                onClose={() => setOpened(false)}
            >
                <InputWrapper label="Customer Name" required>
                    <TextInput
                        placeholder="John Doe"
                        onChange={(e) => setCustomerName(e.target.value)}
                        required
                    />
                </InputWrapper>
                <Space h="md" />
                <InputWrapper label="Customer Address" required>
                    <TextInput
                        placeholder="123 Main St"
                        onChange={(e) => setCustomerAddress(e.target.value)}
                        required
                    />
                </InputWrapper>
                <Space h="md" />
                <Divider />
                <Space h="md" />
                <InputWrapper label="Item Name" required>
                    <TextInput
                        placeholder="Item Name"
                        onChange={(e) => setItemName(e.target.value)}
                        required
                    />
                </InputWrapper>
                <Space h="md" />
                <InputWrapper label="Item Size" required>
                    <Textarea
                        placeholder="Item Size"
                        onChange={(e) => setItemDesc(e.target.value)}
                        required
                    />
                </InputWrapper>
                <Space h="md" />
                <Group grow>
                    <InputWrapper label="Item SKU" required>
                        <TextInput
                            placeholder="Item SKU"
                            onChange={(e) => setItemSku(e.target.value)}
                            required
                        />
                    </InputWrapper>
                    <InputWrapper label="Item Price" required>
                        <TextInput
                            placeholder="Item Price"
                            type="number"
                            onChange={(e) =>
                                setItemPrice(Number(e.target.value))
                            }
                            required
                        />
                    </InputWrapper>
                    <InputWrapper label="Item Quantity" required>
                        <TextInput
                            placeholder="Item Quantity"
                            type="number"
                            onChange={(e) =>
                                setItemQuantity(Number(e.target.value))
                            }
                            required
                        />
                    </InputWrapper>
                </Group>
                <Space h="md" />
                <Group position="right">
                    <Button
                        color="green"
                        onClick={() => {
                            setInvoiceItems([
                                {
                                    name: itemName,
                                    size: itemDesc,
                                    sku: itemSku,
                                    price: itemPrice,
                                    quantity: itemQuantity,
                                },
                                ...invoiceItems,
                            ]);
                            setItemName("");
                            setItemDesc("");
                            setItemSku("");
                            setItemPrice(0);
                            setItemQuantity(0);
                        }}
                    >
                        Add Item
                    </Button>
                </Group>
                <Space h="md" />
                <Table>
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Size</th>
                            <th>SKU</th>
                            <th>$</th>
                            <th>Qty</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {invoiceItems.map((item) => (
                            <InvoiceItemRow
                                key={item.name}
                                item={item}
                                invoiceItems={invoiceItems}
                                setInvoiceItems={setInvoiceItems}
                            />
                        ))}
                    </tbody>
                </Table>
                <Space h="md" />
                <Divider />
                <Space h="md" />
                <Group position="right">
                    <Button color="green" onClick={generateInvoice}>
                        Generate
                    </Button>
                </Group>
            </Modal>
        </>
    );
};

const Invoice = () => {
    const [showInvoiceConfigurator, setShowInvoiceConfigurator] =
        useState(false);

    return (
        <>
            <Box sx={containerStyles}>
                <h3>Invoice Generator</h3>
                <Space h="md" />
                <Button
                    color="blue"
                    onClick={() => setShowInvoiceConfigurator(true)}
                >
                    Configure Invoice
                </Button>
            </Box>
            <InvoiceConfiguratorModal
                opened={showInvoiceConfigurator}
                setOpened={setShowInvoiceConfigurator}
            />
        </>
    );
};

export default Invoice;
