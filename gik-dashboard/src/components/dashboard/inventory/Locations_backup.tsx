// import {
//     Modal,
//     LoadingOverlay,
//     Autocomplete,
//     Group,
//     TextInput,
//     Button,
//     ActionIcon,
//     Space,
//     Text,
//     Box,
//     Table,
//     Checkbox,
// } from "@mantine/core";
// import { showNotification } from "@mantine/notifications";
// import { Dispatch, SetStateAction, useState, useEffect } from "react";
// import QRCode from "react-qr-code";
// import { Trash, Qrcode, CirclePlus } from "tabler-icons-react";
// import { containerStyles } from "../../../styles/container";

// import type { Location } from "../../../types/location";

// import styles from "../../../styles/Location.module.scss";

// interface data {
//     name: string;
//     sku: string;
// }

// export const CreateLocationModal = ({
//     opened,
//     setOpened,
//     refresh,
// }: {
//     opened: boolean;
//     setOpened: Dispatch<SetStateAction<boolean>>;
//     refresh: () => Promise<void>;
// }) => {
//     const [suggestions, setSuggestions] = useState<string[]>([]);

//     const [itemName, setItemName] = useState("");

//     const [locationName, setLocationName] = useState("");
//     const [locationLetter, setLocationLetter] = useState("");

//     const [loading, setLoading] = useState(false);

//     const fetchSuggestions = async () => {
//         const response = await fetch(
//             `${process.env.REACT_APP_API_URL}/items/suggest?query=${itemName}`,
//             {
//                 credentials: "include",
//             }
//         );

//         const data: {
//             success: boolean;
//             message: string;
//             data: data[];
//         } = await response.json();


//         let temp: string[]

//         temp = []

//         if (data.success) {

//             for (let i = 0; i < data.data.length; i++) {
//                 temp = [...temp, data.data[i].name]
//             }
//         }

//         setSuggestions(temp);
//     };

//     useEffect(() => {
//         fetchSuggestions();
//     }, [itemName]);

//     const doCreate = async () => {
//         setItemName("")
//         setLoading(true);

//         const response = await fetch(
//             `${process.env.REACT_APP_API_URL}/location/add`,
//             {
//                 credentials: "include",
//                 method: "PUT",
//                 headers: {
//                     "Content-Type": "application/json",
//                 },
//                 body: JSON.stringify({
//                     name: locationName,
//                     //letter: locationLetter,
//                     productName: itemName,
//                 }),
//             }
//         );

//         setLoading(false);

//         const data: {
//             success: boolean;
//             message: string;
//         } = await response.json();

//         if (data.success) {
//             await refresh();
//             setOpened(false);
//             return;
//         }

//         showNotification({
//             color: "red",
//             title: "Error creating location",
//             message: data.message,
//         });
//     };

//     return (
//         <>
//             <Modal
//                 title="Create Location"
//                 opened={opened}
//                 onClose={() => setOpened(false)}
//             >
//                 <LoadingOverlay visible={loading} />
//                 <Autocomplete
//                     onChange={setItemName}
//                     label="Item Name OR Product ID"
//                     placeholder="Start typing..."
//                     data={suggestions}
//                 />
//                 <Space h="md" />
//                 <Group grow>
//                     <TextInput
//                         label="Name"
//                         placeholder="M15"
//                         onChange={(e) => setLocationName(e.target.value)}
//                     />
//                 </Group>
//                 <Space h="md" />
//                 <Group position="right">
//                     <Button color="green" onClick={doCreate}>
//                         Create
//                     </Button>
//                 </Group>
//             </Modal>
//         </>
//     );
// };

// export const LocationQrGenerator = ({
//     location,
//     opened,
//     setOpened,
// }: {
//     location: string;
//     opened: boolean;
//     setOpened: Dispatch<SetStateAction<boolean>>;
// }) => {
//     const getQrFile = async () => {
//         const response = await fetch(
//             `${process.env.REACT_APP_API_URL}/qr/codes?labels=${location}`,
//             {
//                 credentials: "include",
//             }
//         );

//         // data is arraybuffer
//         const data = await response.arrayBuffer();

//         // convert to blob
//         const blob = new Blob([data], { type: "application/pdf" });

//         // create a url from the blob
//         const url = URL.createObjectURL(blob);

//         // open the url with the pdf viewer
//         window.open(url, "_blank");
//     };

//     return (
//         <>
//             <Modal
//                 opened={opened}
//                 onClose={() => setOpened(false)}
//                 title="Location QR Code Generator"
//             >
//                 <Text>{location}</Text>
//                 <QRCode value={location} size={200} />
//                 <Space h="md" />
//                 <Group position="right">
//                     <Button color="green" onClick={getQrFile}>
//                         Print
//                     </Button>
//                 </Group>
//             </Modal>
//         </>
//     );
// };

// export const LocationComponent = ({
//     location,
//     refresh,
//     locationIdentifiers,
//     setLocationIdentifiers,
// }: {
//     location: Location;
//     refresh: () => Promise<void>;
//     locationIdentifiers: string[];
//     setLocationIdentifiers: Dispatch<SetStateAction<string[]>>;
// }) => {
//     const [showQr, setShowQr] = useState(false);

//     const doDelete = async () => {
//         await fetch(
//             `${process.env.REACT_APP_API_URL}/location/delete?id=${location.ID}`,
//             {
//                 method: "DELETE",
//                 credentials: "include",
//             }
//         );

//         await refresh();
//     };

//     const addSub = async () => {
//         await fetch(
//             `${process.env.REACT_APP_API_URL}/location/add/sub?name=${location.name}`,
//             {
//                 method: "PUT",
//                 credentials: "include",
//             }
//         );

//         refresh();
//     };

//     return (
//         <>
//             <tr className={styles.locationComponent}>
//                 <td>
//                     <Checkbox
//                         onClick={() => {
//                             // check if it's in locationIdentifiers
//                             const identifier = `${location.name}${location.letter}`;

//                             if (locationIdentifiers.includes(identifier)) {
//                                 // remove it
//                                 setLocationIdentifiers(
//                                     locationIdentifiers.filter(
//                                         (i) => i !== identifier
//                                     )
//                                 );
//                                 return;
//                             }

//                             setLocationIdentifiers([
//                                 ...locationIdentifiers,
//                                 identifier,
//                             ]);
//                         }}
//                     />
//                 </td>
//                 <td>{location.name+location.letter}</td>
//                 <td>{location.productName}</td>
//                 <td>
//                     <Group>
//                         <ActionIcon variant="default" onClick={doDelete}>
//                             <Trash />
//                         </ActionIcon>
//                         <ActionIcon
//                             variant="default"
//                             onClick={() => setShowQr(true)}
//                         >
//                             <Qrcode />
//                         </ActionIcon>
//                         <ActionIcon
//                             variant="default"
//                             onClick={addSub}
//                         >
//                             <CirclePlus />
//                         </ActionIcon>
//                     </Group>
//                 </td>
//             </tr>
//             <LocationQrGenerator
//                 opened={showQr}
//                 setOpened={setShowQr}
//                 location={location.name + location.letter}
//             />
//         </>
//     );
// };

// export const LocationsManager = () => {
//     const [locations, setLocations] = useState<Location[]>([]);

//     const [nameTyping, setNameTyping] = useState("");
//     const [filterName, setFilterName] = useState("");

//     const [letterTyping, setLetterTyping] = useState("");
//     const [filterLetter, setFilterLetter] = useState("");

//     const [productTyping, setProductTyping] = useState("");
//     const [filterProduct, setFilterProduct] = useState("");

//     const [showCreateModal, setShowCreateModal] = useState(false);

//     const [locationIdentifiers, setLocationIdentifiers] = useState<string[]>(
//         []
//     );

//     const bulkPrint = async () => {
//         const response = await fetch(
//             `${
//                 process.env.REACT_APP_API_URL
//             }/qr/codes?labels=${encodeURIComponent(
//                 locationIdentifiers.join(" ")
//             )}`,
//             {
//                 credentials: "include",
//             }
//         );

//         // data is arraybuffer
//         const data = await response.arrayBuffer();

//         // convert to blob
//         const blob = new Blob([data], { type: "application/pdf" });

//         // create a url from the blob
//         const url = URL.createObjectURL(blob);

//         // open the url with the pdf viewer
//         window.open(url, "_blank");
//     };

//     const fetchLocations = async () => {
//         const response = await fetch(
//             `${process.env.REACT_APP_API_URL}/location/list?name=${filterName}&letter=${filterLetter}&product=${filterProduct}`,
//             {
//                 credentials: "include",
//             }
//         );

//         const data: {
//             success: boolean;
//             message: string;
//             data: Location[];
//         } = await response.json();

//         if (data.success) {
//             setLocations(data.data);
//         }
//     };

//     useEffect(() => {
//         fetchLocations();
//     }, [filterName, filterLetter, filterProduct]);

//     return (
//         <>
//             <Box sx={containerStyles}>
//                 <Group position="apart">
//                     <h3>Locations</h3>
//                     <ActionIcon
//                         sx={{
//                             height: "4rem",
//                             width: "4rem",
//                         }}
//                         onClick={() => setShowCreateModal(true)}
//                     >
//                         <CirclePlus />
//                     </ActionIcon>
//                 </Group>
//                 <Space h="md" />
//                 <Text>Filter</Text>
//                 <Group>
//                     <TextInput
//                         placeholder="W15"
//                         onChange={(e: any) => setNameTyping(e.target.value)}
//                     />
//                     <TextInput
//                         placeholder="A"
//                         onChange={(e: any) => setLetterTyping(e.target.value)}
//                     />
//                     <TextInput
//                         placeholder="pants OR 18830"
//                         onChange={(e: any) => setProductTyping(e.target.value)}
//                     />
//                     <Button
//                         onClick={() => {
//                             setFilterName(nameTyping);
//                             setFilterLetter(letterTyping);
//                             setFilterProduct(productTyping);
//                         }}
//                     >
//                         Filter
//                     </Button>
//                 </Group>
//                 <Space h="md" />
//                 {locationIdentifiers.length ? (
//                     <Button color="blue" onClick={bulkPrint}>
//                         Bulk Print QR Codes
//                     </Button>
//                 ) : null}
//                 <Space h="md" />
//                 <Table>
//                     <thead>
//                         <tr>
//                             <th>[+]</th>
//                             <th>Name</th>
//                             <th>Product Name</th>
//                             <th>Actions</th>
//                         </tr>
//                     </thead>
//                     <tbody>
//                         {locations.map((location) => (
//                             <LocationComponent
//                                 key={location.ID}
//                                 location={location}
//                                 refresh={fetchLocations}
//                                 locationIdentifiers={locationIdentifiers}
//                                 setLocationIdentifiers={setLocationIdentifiers}
//                             />
//                         ))}
//                     </tbody>
//                 </Table>
//                 <Group position="right">
//                     <Button onClick={fetchLocations}>
//                         Refresh
//                     </Button>
//                 </Group>
//             </Box>
//             <CreateLocationModal
//                 refresh={fetchLocations}
//                 opened={showCreateModal}
//                 setOpened={setShowCreateModal}
//             />
//         </>
//     );
// };
