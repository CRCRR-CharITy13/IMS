import {
    Box,
    Center,
    Group,
    Pagination,
    Skeleton,
    Table,
    Text,
} from "@mantine/core";

import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from "chart.js";
import { useEffect, useState } from "react";
import { Line } from "react-chartjs-2";
import { Item } from "../../types/item";
import { AdvancedLog } from "../../types/logs";

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

const ItemRow = ({ item }: { item: Item }) => {
    return (
        <>
            <tr id={String(item.ID)}>
                <td>{item.name}</td>
                <td>{item.sku || "None"}</td>
                <td>{item.size}</td>
                <td>{item.quantity}</td>
            </tr>
        </>
    );
};

const last7Days = () => {
    return "0123456"
        .split("")
        .map(function (n) {
            const d = new Date();
            d.setDate(d.getDate() - parseInt(n));

            return (function (day, month, year) {
                return [
                    day < 10 ? "0" + day : day,
                    month < 10 ? "0" + month : month,
                    year,
                ].join("/");
            })(d.getDate(), d.getMonth(), d.getFullYear());
        })
        .reverse()
        .join(",");
};

const options = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
        legend: {
            display: false,
        },
    },
    scales: {
        x: {
            ticks: {
                autoSkip: false,
                minRotation: 20,
            },
        },
        /*y: {
            grid: {
                drawBorder: false,
                color: (context) => {
                    if (context.tick.value != 0) {
                        return;
                    }

                    return '#000000';
                },
            },
        }*/
    },
};


const TrendingItems = () => {
    const [items, setItems] = useState<Item[]>([]);

    const getTrendingItems = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/analytics/trending`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: Item[];
        } = await response.json();

        if (data.success) {
            setItems(data.data);
        }
    };

    useEffect(() => {
        getTrendingItems();
    }, []);

    return (
        <>
            {" "}
            <Box
                sx={{
                    width: "25rem",
                    height: "20rem",
                    borderRadius: "15px",
                    boxShadow: ` 5px 5px 10px var(--neumorphism),
            -5px -5px 10px var(--inverted-text)`,
                    padding: "1rem",
                    display: "flex",
                    flexDirection: "column",
                    flexGrow: 1,
                }}
            >
                <h2>Trending Items</h2>
                <Skeleton
                    height={"90%"}
                    width={"100%"}
                    visible={items.length === 0}
                >
                    <Table>
                        <thead>
                            <tr>
                                <th>Name</th>
                                <th>SKU</th>
                                <th>Size</th>
                                <th>Stock</th>
                            </tr>
                        </thead>
                        <tbody>
                            {items.map((item) => (
                                <ItemRow key={item.ID} item={item} />
                            ))}
                        </tbody>
                    </Table>
                </Skeleton>
            </Box>
        </>
    );
};

const skeletonStyles = {
    "::before": { background: "var(--inverted-text)" },
    "::after": { background: "var(--skeleafter)" },
    overflowY: "scroll",
};

const AttentionRequired = () => {
    const [attentionData, setAttentionData] = useState<Item[]>([]);
    const [attentionLoading, setAttentionLoading] = useState<boolean>(true);

    const [currentPage, setCurrentPage] = useState<number>(1);
    const [totalPages, setTotalPages] = useState<number>(1);

    useEffect(() => {
        fetchDataAttention();
    }, [currentPage]);

    const fetchDataAttention = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/analytics/attention?page=${currentPage}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            message: string;
            data: {
                data: Item[];
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setAttentionData(data.data.data);

            setTotalPages(data.data.totalPages);

            setAttentionLoading(false);
        }
    };

    return (
        <>
            <Box
                sx={{
                    width: "25rem",
                    height: "20rem",
                    borderRadius: "15px",
                    padding: "1rem",
                    display: "flex",
                    flexDirection: "column",
                    flexGrow: 1,
                }}
            >
                <h2>Attention Required</h2>
                <Skeleton
                    height={"90%"}
                    width={"100%"}
                    visible={attentionLoading}
                >
                    <Table>
                        <thead>
                            <tr>
                                <th>Name</th>
                                <th>SKU</th>
                                <th>Size</th>
                                <th>Stock</th>
                            </tr>
                        </thead>
                        <tbody>
                            {attentionData.map((item) => (
                                <ItemRow key={item.ID} item={item} />
                            ))}
                        </tbody>
                    </Table>
                    <Center>
                        <Pagination
                            total={totalPages}
                            value={currentPage}
                            onChange={setCurrentPage}
                        />
                    </Center>
                </Skeleton>
            </Box>
        </>
    );
};

const RecentActivityLog = ({ log }: { log: AdvancedLog }) => {
    const [username, setUsername] = useState<string>("");

    const getUsername = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/info/username?id=${log.userId}`,
            {
                credentials: "include",
            }
        );

        const data = await response.json();

        setUsername(data.data);
    };

    useEffect(() => {
        getUsername();
    }, []);

    return (
        <>
            <Box
                sx={{
                    padding: "1rem",
                    boxSizing: "border-box",
                    borderRadius: "10px",
                    userSelect: "none",
                    "&:hover": {
                        backgroundColor: "#d9d9d9",
                    },
                }}
            >
                <Text>
                    <b>{username}</b> {log.action.toLowerCase()} at{" "}
                    <b>{new Date(log.timestamp * 1000).toLocaleString()}</b>.
                </Text>
            </Box>
        </>
    );
};

const RecentActivity = () => {
    const [logs, setLogs] = useState<AdvancedLog[]>([]);

    useEffect(() => {
        getActivity();
    }, []);

    const getActivity = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/analytics/activity`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            message: string;
            data: AdvancedLog[];
        } = await response.json();

        if (data.success) {
            setLogs(data.data);
        }
    };

    return (
        <>
            <Box
                sx={{
                    width: "25rem",
                    height: "20rem",
                    borderRadius: "15px",
                    padding: "1rem",
                    display: "flex",
                    flexDirection: "column",
                    flexGrow: 1,
                }}
            >
                <h2>Recent Activity</h2>
                <Skeleton
                    height={"90%"}
                    width={"100%"}
                    visible={logs.length === 0}
                >
                    {logs.map((log) => (
                        <RecentActivityLog key={log.ID} log={log} />
                    ))}
                </Skeleton>
            </Box>
        </>
    );
};

const Analytics = () => {
    const [importData, setImportData] = useState<number[]>([]);
    const [importLabels, setImportLabels] = useState<string[]>([]);
    const [importLoading, setImportLoading] = useState<boolean>(true);

    const [exportData, setExportData] = useState<number[]>([]);
    const [exportLabels, setExportLabels] = useState<string[]>([]);
    const [exportLoading, setExportLoading] = useState<boolean>(true);

    const [totalStockData, setTotalStockData] = useState<number[]>([]);
    const [totalStockLabels, setTotalStockLabels] = useState<string[]>([]);
    const [totalStockLoading, setTotalStockLoading] = useState<boolean>(true);

    useEffect(() => {
        init();
    }, []);

    const init = async () => {
        // Temporary commented by tuan on 22 Sept 2023, to avoid querying transaction
        // await fetchDataImport();
        // await fetchDataExport();
        // await fetchDataTotalStock();
    };

    const fetchDataImport = async () => {
        const responseImport = await fetch(
            `${process.env.REACT_APP_API_URL}/analytics/transaction?type=true`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: [];
        } = await responseImport.json();

        
        if (data.success) {
            setImportData(data.data);

            const labels: string[] = last7Days().split(",");

            setImportLabels(labels);

            setImportLoading(false);
        }
    };

    const fetchDataExport = async () => {
        const responseImport = await fetch(
            `${process.env.REACT_APP_API_URL}/analytics/transaction?type=false`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: [];
        } = await responseImport.json();

        if (data.success) {
            setExportData(data.data);

            const labels: string[] = last7Days().split(",");

            setExportLabels(labels);

            setExportLoading(false);
        }
    };

    const fetchDataTotalStock = async () => {
        const responseImport = await fetch(
            `${process.env.REACT_APP_API_URL}/analytics/transaction/total`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: [];
        } = await responseImport.json();

        if (data.success) {
            setTotalStockData(data.data);

            const labels: string[] = last7Days().split(",");

            setTotalStockLabels(labels);

            setTotalStockLoading(false);
        }
    };

    return (
        <>
            <Box
                sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    flexWrap: "wrap",
                    gap: "1.5rem",
                }}
            >
                <Box
                    sx={{
                        width: "25rem",
                        height: "20rem",
                        borderRadius: "15px",
                        padding: "1rem",
                        display: "flex",
                        flexDirection: "column",
                        flexGrow: 1,
                    }}
                >
                    <h2>Net Stock Change</h2>
                    <Skeleton
                        height={"90%"}
                        width={"100%"}
                        visible={totalStockLoading}
                    >
                        {importData.length && (
                            <Line
                                options={options}
                                data={{
                                    labels: totalStockLabels,
                                    datasets: [
                                        {
                                            label: "Daily Stock Change",
                                            data: totalStockData,
                                            tension: 0.25,
                                            borderColor: "#009f0b",
                                        },
                                    ],
                                }}
                            />
                        )}
                    </Skeleton>
                </Box>

                <Box
                    sx={{
                        width: "30rem",
                        height: "20rem",
                        borderRadius: "15px",
                        padding: "1rem",
                        display: "flex",
                        flexDirection: "column",
                        flexGrow: 1,
                    }}
                >
                    <h2>Imports</h2>
                    <Skeleton
                        height={"90%"}
                        width={"100%"}
                        visible={importLoading}
                    >
                        {importData.length && (
                            <Line
                                options={options}
                                data={{
                                    labels: importLabels,
                                    datasets: [
                                        {
                                            label: "Daily Imports",
                                            data: importData,
                                            fill: false,
                                            borderColor: "rgb(137, 172, 255)",
                                            backgroundColor:
                                                "rgba(32, 99, 255)",
                                            tension: 0.25,
                                        },
                                    ],
                                }}
                            />
                        )}
                    </Skeleton>
                </Box>
                <Box
                    sx={{
                        width: "30rem",
                        height: "20rem",
                        borderRadius: "15px",
                        padding: "1rem",
                        display: "flex",
                        flexDirection: "column",
                        flexGrow: 1,
                    }}
                >
                    <h2>Exports</h2>
                    <Skeleton
                        height={"90%"}
                        width={"100%"}
                        visible={exportLoading}
                    >
                        {importData.length && (
                            <Line
                                options={options}
                                data={{
                                    labels: exportLabels,
                                    datasets: [
                                        {
                                            label: "Daily Exports",
                                            data: exportData,
                                            fill: false,
                                            borderColor: "rgb(255, 143, 167)",
                                            backgroundColor:
                                                "rgba(255, 99, 132, 0.5)",
                                            tension: 0.25,
                                        },
                                    ],
                                }}
                            />
                        )}
                    </Skeleton>
                </Box>
                {/* TODO: Change attention required
                <AttentionRequired /> 
                */}
                <RecentActivity />
            </Box>
        </>
    );
};

export default Analytics;
