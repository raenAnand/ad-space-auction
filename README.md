# Ad Space Auction 
Ad Space Auction is a system that facilitates the auctioning of ad spaces to potential bidders. It consists of two main components: the supply side and the demand side. The supply side manages the ad spaces available for auction, allows creation and updates of ad spaces, and provides APIs to retrieve information about the ad spaces. The demand side handles the bidding process, where bidders can place bids on ad spaces during specific auction periods. 

## Technologies Used
- Go programming language 
- MySQL database 
- Docker 

## Prerequisites 
 - Docker installed on your machine
 - MySQL server running or a remote MySQL server accessible

## Installation
 1.  Clone the repository:

    git clone https://github.com/raenAnand/ad-space-auction.git

 
2. Set up the database:

- Update the database connection details in `supply/database/mysql.go` and `demand/database/mysql.go` with your MySQL server credentials.
- Create the necessary database tables by running the SQL scripts located in `supply/database/schema.sql` and `demand/database/schema.sql`.

3. Build and run the Docker containers:

```bash
cd ad-space-auction
docker-compose up -d
```
 4.  The supply side service should be accessible at `http://localhost:8000`. The demand side service should be accessible at `http://localhost:9000`. The auction side service should be accessible at `http://localhost:7000`.


## Information Architecture:

1.  Supply Side Service:
    
    -   This service is responsible for listing all available ad spaces and their base prices.
    -   It handles the creation, update, and retrieval of ad spaces.
    -   It interacts with the database to store and retrieve ad space information.
    
    API Endpoints:
    
    -   `GET /ad-spaces`: Retrieve a list of all available ad spaces.
    -   `GET /ad-spaces/{id}`: Retrieve details of a specific ad space.
    -   `POST /ad-spaces`: Create a new ad space.
    -   `PUT /ad-spaces/{id}`: Update the details of an existing ad space.
    
    Database Schema:
    
    -   `ad_space` table: id, name, description, base_price
2.  Demand Side Service:
    
    -   This service is responsible for listing all the bidders interested in bidding for the ad spaces.
    -   It handles the creation, update, and retrieval of bidders' information.
    -   It interacts with the database to store and retrieve bidder information.
    
    API Endpoints:
    
    -   `GET /bidders`: Retrieve a list of all bidders.
    -   `GET /bidders/{id}`: Retrieve details of a specific bidder.
    -   `POST /bidders`: Create a new bidder.
    -   `PUT /bidders/{id}`: Update the details of an existing bidder.
    
    Database Schema:
    
    -   `bidder` table: id, name, email, phone
3.  Auction Service:
    
    -   This service manages the auction process for ad spaces.
    -   It handles starting and ending auctions, accepting bids, and selecting the winning bidder.
    -   It interacts with the supply side and demand side services to retrieve necessary information.
    
    API Endpoints:
    
    -   `POST /auctions`: Start a new auction for a specific ad space.
    -   `POST /auctions/{id}/bids`: Accept a bid for a specific auction.
    
    Database Schema:
    
    -   `auction` table: id, ad_space_id, start_time, end_time, winning_bidder_id
    -   `bid` table: id, auction_id, bidder_id, amount

###  Constraints/Assumptions:

-   Each auction has a definite end time, which should be specified during auction creation.
-   Bidding begins as soon as an auction is published.
-   Bids should be accepted only within the specified auction's start and end time.
-   Bidders cannot bid after the auction ends.
-   Only registered bidders can participate in auctions.
-   Bids must be higher than the previous highest bid for a specific auction.
-   Ad spaces can have different base prices.
-   Winning bidder is determined based on the highest bid amount.