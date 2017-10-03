# Problem Statement:- 
In this, we’ll deal with one optimization problem, which can be informally defined as: 
Assume that we have N users and N products (dresses).For each pair (user, product) we have to assign some weight to the dresses and find the most relevant dresses for the user based on the traits and interest of the user. 

Input:- 
A.Product Json:- here Attributes related to product 
B.User Json:- here Attributes related to User behavior and interest 

Output:- Most relevant products to the user based on his/her choice.
Eg:- Sheela fair coloured tall girl of age 26 wants to buy a red dress for her college party. 

So we have to find from our product catalogue the most relevant red dress for her based on her interest like party type, western, 25 yrs, fair colored tall. 

Note:- weight to the attributes should be configurable, either it can be hard-coded in the system or ask the user for the weights.

# Solution:-

According to the problem statement, I have implemented Go based rest apis for solving the problem.

# Running the application

    -> Go to the main folder of the package structure and run the server.go file to boot up the application.
        
        go run server.go

        Server is setup at localhost:3000


Endpoints (with the JSON) to send are:->


GET /users      
    => get all the registered users
            --> No JSON required.. Free API. Gives you JSON array of users 

POST /users     
    => to add a user into the system
            --> JSON =>
            {
                "name": "Puneet", 
                "gender": "male", 
                "age": 22, 
                "height": 172, 
                "color": "fair", 
                "interest": "Party"
            }


GET /users/{id} 
    => to get a particular user by id
            --> No JSON required. Just pass the id in path parameter

DELETE /users/{id}  
    => to delete a user from system
            --> No JSON required. Just pass the id in path parameter

GET /products      => get all the registered products
    --> No JSON required.. Free API. Gives you JSON array of products 

POST /products    
    => to add a product into the system
            --> JSON =>
            {
                "productName": "shirt", 
                "color": "blue", 
                "gender": "male", 
                "ageGroup": "20", 
                "size": "M", 
                "eventType": "casual", 
                "traditionType": "western"
            }

GET /products/{id} 
    => to get a particular product by id
            --> No JSON required. Just pass the id in path parameter

DELETE /products/{id}  
    => to delete a product from system
            --> No JSON required. Just pass the id in path parameter

POST /weights     
    => to configure the weights for the system
             --> JSON =>
                    {
                        "color": 4, 
                        "gender": 6, 
                        "ageGroup": 5, 
                        "size": 3, 
                        "eventType": 2, 
                        "traditionType": 1
                    }

POST /matches     
    => to get the matching product for the system
            --> JSON =>
                {
                    "product" :
                        {
                            "productName": "shirt",
                            "color": "blue", 
                            "size": "M", 
                            "eventType": "casual", 
                            "traditionType": "western"
                        },
	                "user" : 
                        {
		                    "name": "Savita", 
                            "gender": "female", 
                            "age": 25, 
                            "height": 172, 
                            "color": "Blue", 
                            "interest": "Party"
	                    }
                }

    => OUTPUT is the JSON array of the matching products in sorted order 


# Data for application

I have used MongoDb for storing the data to make it persistent. 
However, I have not used any caching for optimizing the database call. 
I am just getting the whole data from the mongo at once and then performing the match operation by the various filters within the code.
You can add data into the application as per requirement and it will be stored in the mongodb itself.

I have provided some sample data to perform test the APIs with application.

products.json

users.json

weights.json
You can import these files into your local database using the commands:

    mongoimport -d eCommerce -c products --file products.json
    
    mongoimport -d eCommerce -c users --file users.json
    
    mongoimport -d eCommerce -c weights --file weights.json

Also you can use any of the APIs to add users data and products data or configure the weights for the attributes.
