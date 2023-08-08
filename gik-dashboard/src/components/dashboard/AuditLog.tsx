import {
    Box,
    Button,
    Center,
    Input,
    InputWrapper,
    Table,
    SegmentedControl,
    Accordion,
    Tab,
    Pagination,
    Space,
    LoadingOverlay,
} from "@mantine/core";

import { DateRangePicker } from "@mantine/dates";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { AdvancedLog } from "../../types/logs";

import styles from "../../styles/AuditLog.module.scss";

const exampleData = [
    {
        id: 1,
        date: "2022-07-04",
        user: "amy",
        action: "Added 10 credits to joe",
    },
    {
        id: 2,
        date: "2022-07-04",
        user: "joe",
        action: "Added 1,000,000 credits to joe",
    },
    {
        id: 3,
        date: "2022-07-04",
        user: "amy",
        action: "Subtracted 1,000,000 credits from joe",
    },
    {
        id: 4,
        date: "2022-07-04",
        user: "amy",
        action: "Deleted the entire database",
    },
];

const AdvancedLogs = ({
    actionFilter,
    dateFilter,
    userFilter,
    setVisible,
}: {
    actionFilter: string;
    dateFilter: [Date | null, Date | null] | undefined;
    setVisible: Dispatch<SetStateAction<boolean>>;
    userFilter: string;
}) => {
    const [data, setData] = useState<AdvancedLog[]>([]);

    const [currentPage, setCurrentPage] = useState<number>(1);

    const [totalPages, setTotalPages] = useState<number>(1);

    const getData = async () => {
        setVisible(true);

        const response = await fetch(
            `${
                process.env.REACT_APP_API_URL
            }/logs/advanced?page=${currentPage}&action=${actionFilter}&date=${
                dateFilter && dateFilter[0] && dateFilter[1]
                    ? `${dateFilter[0].getTime() / 1000},${
                          dateFilter[1].getTime() / 1000
                      }`
                    : ""
            }&user=${userFilter}`,
            {
                credentials: "include",
            }
        );

        setVisible(false);

        const data: {
            success: boolean;
            data: {
                data: AdvancedLog[];
                total: number;
                currentPage: number;
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setData(data.data.data);
            setTotalPages(data.data.totalPages);
        }
    };

    useEffect(() => {
        getData();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        getData();
    }, [actionFilter, dateFilter, userFilter]);

    return (
        <>
            <Accordion multiple>
                {data.map((log) => (
                    <Accordion.Item label={`${log.path}`}>
                        {Object.entries(log).map((entry) => {
                            const key = entry[0];
                            const value = entry[1];

                            const excludedKeys = [
                                "CreatedAt",
                                "UpdatedAt",
                                "DeletedAt",
                            ];

                            if (excludedKeys.includes(key)) return;

                            return (
                                <>
                                    <div className={styles.log}>
                                        <p>
                                            <b>{key}</b>:{" "}
                                            {key === "timestamp" ? (
                                                <span>
                                                    {new Date(
                                                        value * 1000
                                                    ).toLocaleString()}
                                                </span>
                                            ) : (
                                                <span>{value}</span>
                                            )}
                                        </p>
                                    </div>
                                </>
                            );
                        })}
                    </Accordion.Item>
                ))}
            </Accordion>
            <Center
                sx={{
                    marginTop: "1rem",
                }}
            >
                <Pagination
                    page={currentPage}
                    onChange={setCurrentPage}
                    total={totalPages}
                />
            </Center>
        </>
    );
};

const SimpleLogRow = ({ log }: { log: AdvancedLog }) => {
    const [username, setUsername] = useState<string>("");

    const getUsername = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/info/username?id=${log.userId}`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: string;
        } = await response.json();

        if (data.success) {
            setUsername(data.data);
        }
    };

    useEffect(() => {
        getUsername();
    }, []);

    return (
        <>
            <tr>
                <td>{log.ID}</td>
                <td>{username}</td>
                <td>{log.ipAddress}</td>
                <td>{log.action}</td>
            </tr>
        </>
    );
};

const SimpleLogs = ({
    actionFilter,
    dateFilter,
    userFilter,
    setVisible,
}: {
    actionFilter: string;
    dateFilter: [Date | null, Date | null] | undefined;
    setVisible: Dispatch<SetStateAction<boolean>>;
    userFilter: string;
}) => {
    const [data, setData] = useState<AdvancedLog[]>([]);

    const [currentPage, setCurrentPage] = useState<number>(1);

    const [totalPages, setTotalPages] = useState<number>(1);

    const getData = async () => {
        setVisible(true);

        const response = await fetch(
            `${
                process.env.REACT_APP_API_URL
            }/logs/simple?page=${currentPage}&action=${actionFilter}&date=${
                dateFilter && dateFilter[0] && dateFilter[1]
                    ? `${dateFilter[0].getTime() / 1000},${
                          dateFilter[1].getTime() / 1000
                      }`
                    : ""
            }&user=${userFilter}`,
            {
                credentials: "include",
            }
        );

        setVisible(false);

        const data: {
            success: boolean;
            data: {
                data: AdvancedLog[];
                total: number;
                currentPage: number;
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setData(data.data.data);
            setTotalPages(data.data.totalPages);
        }
    };

    useEffect(() => {
        getData();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        getData();
    }, [actionFilter, dateFilter, userFilter]);

    return (
        <>
            <Table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Username</th>
                        <th>IP</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((log) => (
                        <SimpleLogRow key={log.ID} log={log} />
                    ))}
                </tbody>
            </Table>
            <Center
                sx={{
                    marginTop: "1rem",
                }}
            >
                <Pagination
                    page={currentPage}
                    onChange={setCurrentPage}
                    total={totalPages}
                />
            </Center>
        </>
    );
};

const AuditLog = () => {
    const [viewMode, setViewMode] = useState<"smp" | "adv">("smp");

    const [actionFilter, setActionFilter] = useState<string>("");
    const [actionFilterEditing, setActionFilterEditing] = useState<string>("");

    const [dateFilter, setDateFilter] = useState<
        [Date | null, Date | null] | undefined
    >();
    const [dateFilterEditing, setDateFilterEditing] = useState<
        [Date | null, Date | null] | undefined
    >();

    const [userFilter, setUserFilter] = useState<string>("");
    const [userFilterEditing, setUserFilterEditing] = useState<string>("");

    const [visible, setVisible] = useState<boolean>(false);

    const doFilter = async () => {
        setActionFilter(actionFilterEditing);
        setDateFilter(dateFilterEditing);
        setUserFilter(userFilterEditing);
    };

    return (
        <>
            <div
                style={{
                    display: "flex",
                    width: "100%",
                    justifyContent: "flex-start",
                    alignItems: "center",
                    flexDirection: "column",
                    height: "100%",
                }}
            >
                <LoadingOverlay visible={visible} />
                <Box
                    sx={{
                        borderRadius: "5px",
                        border: "1px solid gray",
                        width: "100%",
                        padding: "1rem",
                        boxSizing: "border-box",
                        backgroundColor: "var(--inverted-text)",
                        "@media only screen and (max-width: 800px)": {
                            width: "100%",
                        },
                    }}
                >
                    <SegmentedControl
                        data={[
                            { label: "Simple", value: "smp" },
                            {
                                label: "Advanced",
                                value: "adv",
                            },
                        ]}
                        sx={{
                            marginBottom: "1rem",
                        }}
                        onChange={(value: "smp" | "adv") => setViewMode(value)}
                    />

                    <h1>
                        {viewMode === "smp" ? "Simple" : "Advanced"} Audit Logs
                    </h1>
                    <h2>View and filter actions and transactions.</h2>
                    <Center
                        sx={{
                            justifyContent: "space-between",
                            gap: ".5rem",
                            alignItems: "flex-end",
                            marginBottom: "1rem",
                        }}
                    >
                        <DateRangePicker
                            placeholder="Pick Date Range"
                            label="Date Range"
                            onChange={setDateFilterEditing}
                        />
                        <InputWrapper label="User">
                            <Input
                                sx={{
                                    display: "flex",
                                    flexGrow: 1,
                                }}
                                placeholder="amy"
                                onChange={(e: any) =>
                                    setUserFilterEditing(e.target.value)
                                }
                            />
                        </InputWrapper>
                        <InputWrapper label="Action">
                            <Input
                                placeholder="add"
                                sx={{
                                    display: "flex",
                                    flexGrow: 1,
                                }}
                                onChange={(e: any) =>
                                    setActionFilterEditing(e.target.value)
                                }
                            />
                        </InputWrapper>
                        <Button color="green" onClick={doFilter}>
                            Filter
                        </Button>
                    </Center>
                    {/* <Table>
                        {" "}
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Date</th>
                                <th>User</th>
                                <th>Action</th>
                            </tr>
                        </thead>
                        <tbody>{rows}</tbody>
                    </Table> */}
                    {/* <Accordion> */}
                    {viewMode === "smp" ? (
                        <SimpleLogs
                            actionFilter={actionFilter}
                            dateFilter={dateFilter}
                            setVisible={setVisible}
                            userFilter={userFilter}
                        />
                    ) : (
                        <AdvancedLogs
                            actionFilter={actionFilter}
                            dateFilter={dateFilter}
                            setVisible={setVisible}
                            userFilter={userFilter}
                        />
                    )}
                    {/* </Accordion> */}
                </Box>
            </div>
        </>
    );
};

export default AuditLog;
