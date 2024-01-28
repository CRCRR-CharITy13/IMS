import {
    Box,
    Center,
    Group,
    Pagination,
    Skeleton,
    Table,
    Text,
    LoadingOverlay,
    SegmentedControl,
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
import { AdvancedLog } from "../../types/logs";
import { AttentionItem, AttentionLocation } from "../../types/attention";
import { containerStyles } from "../../styles/container";


ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

const AttentionItemRow = ({ attentionItem } : { attentionItem : AttentionItem }) => {
    return (
        <>
            <tr id = {String(attentionItem.ID)}>
                <td>{attentionItem.ID}</td>
                <td>{attentionItem.sku}</td>
                <td>{attentionItem.name}</td>
                <td>{attentionItem.size}</td>
                <td>{attentionItem.message}</td>
            </tr>
        </>
    );
};

const AttentionLocationRow = ({ attentionLocation } : { attentionLocation : AttentionLocation }) => {
    return (
        <>
            <tr id = {String(attentionLocation.ID)}>
                <td>{attentionLocation.ID}</td>
                <td>{attentionLocation.name}</td>
                <td>{attentionLocation.description}</td>
                <td>{attentionLocation.message}</td>
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



const AttentionItemRequired = () => {
    const [attentionItemData, setAttentionItemData] = useState<AttentionItem[]>([]);
    const [attentionItemLoading, setAttentionItemLoading] = useState<boolean>(true);

    const [currentPage, setCurrentPage] = useState<number>(1);
    const [totalPages, setTotalPages] = useState<number>(1);

    useEffect(() => {
        fetchDataAttentionItem();
    }, [currentPage]);

    const fetchDataAttentionItem = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/analytics/attention-item?page=${currentPage}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            message: string;
            data: {
                data: AttentionItem[];
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setAttentionItemData(data.data.data);
            setTotalPages(data.data.totalPages);
            setAttentionItemLoading(false);
        }
    };

    return (
        <>
            
                    <Table>
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>SKU</th>
                                <th>Name</th>
                                <th>Size</th>
                                <th>Message</th>
                            </tr>
                        </thead>
                        <tbody>
                            {attentionItemData.map((attentionItem) => (
                                <AttentionItemRow key={attentionItem.ID} attentionItem={attentionItem} />
                            ))}
                        </tbody>
                    </Table>
            
        </>
    );
};

const AttentionLocationRequired = () => {
    const [attentionLocationData, setAttentionLocationData] = useState<AttentionLocation[]>([]);
    const [attentionLocationLoading, setAttentionLocationLoading] = useState<boolean>(true);

    const [currentPage, setCurrentPage] = useState<number>(1);
    const [totalPages, setTotalPages] = useState<number>(1);

    useEffect(() => {
        fetchDataAttentionLocation();
    }, [currentPage]);

    const fetchDataAttentionLocation = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/analytics/attention-location?page=${currentPage}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            message: string;
            data: {
                data: AttentionLocation[];
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setAttentionLocationData(data.data.data);
            setTotalPages(data.data.totalPages);
            setAttentionLocationLoading(false);
        }
    };

    return (
        <>
            <Table>
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Description</th>
                        <th>Message</th>
                    </tr>
                </thead>
                <tbody>
                    {attentionLocationData.map((attentionLocation) => (
                        <AttentionLocationRow key={attentionLocation.ID} attentionLocation={attentionLocation} />
                    ))}
                </tbody>
            </Table>
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
    const [visible, setVisible] = useState<boolean>(false);
    const [viewMode, setViewMode] = useState<"Items" | "Locations">("Items");


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
                {/* <Box
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
                </Box> */}

                {/* <Box
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
                </Box> */}
                {/* <Box
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
                </Box> */}
                
                <div
                style={containerStyles}>
                <LoadingOverlay visible={visible} />
                <Box sx={containerStyles}>
                    <h2>Attentions</h2>
                    <SegmentedControl
                        data={[
                                { label: "Items", value: "Items" },
                                { label: "Locations", value: "Locations" },
                        ]}
                        sx={{
                            marginBottom: "1rem",
                        }}
                        onChange={(value: "Items" | "Locations") => setViewMode(value)}
                    />

                    <h3>
                        {viewMode === "Items" ? "Items" : "Locations"} Required
                    </h3>
                    
                    {viewMode === "Items" ? (
                        <AttentionItemRequired/>
                    ) : (
                        <AttentionLocationRequired/>
                    )}
                    
                </Box>
            </div>
                <RecentActivity />
            </Box>
        </>
    );
};

export default Analytics;
