# IMS
A web-based Inventory Management System for the Gifts in Kind Program built by CharITy13 built using Golang and React.

The goal of this application is to replace the currently used Google Spreadsheet system with an easier to use and less error prone system that is intuiative and can be used going forward.

## Features
**Inventory Management System**: Manage stocks and items on an on-going basis.

**Warehouse Management System**: Have multiple locations and manage where individual items are.

**Transaction Management System**: Create and keep track of transactions, making the required changes to the stock and locations. Create and keep track of clients and donors and their balances.

**Security**: The system uses authentication in order to utilize it, with the ability to create new users for volunteers and others who require access.

## Project Layout

    IMS/
    ├─ gik-api/                       the back-end
    │  ├─ assets                         semi-permenant data files
    │  ├─ database                       code to run the database
    │  ├─ env                            code to manage the environment variables
    │  ├─ src/                           source code
    │  │  ├─ middleware                      middleware code
    │  │  ├─ routers                         router + endpoints
    │  ├─ type_news                      data structure
    │  ├─ types                          data structures (old)
    │  ├─ utils                          collection of functions
    ├─ gik-dashboard/                 the front-end
    │  ├─ public                         images and files for the front-end
    │  ├─ src/                           source code
    │  │  ├─ assets                        additional images and files
    │  │  ├─ components                    individual panes of the dashboard
    │  │  ├─ routes                        the primary routes of the front-end
    │  │  ├─ styles                        styling sheets
    │  │  ├─ types                         data structures

## How to install

Please see [Getting Set Up](../../wiki/Getting-Set-Up).
