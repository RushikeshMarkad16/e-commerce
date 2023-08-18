# e-commerce

### Problem Statement : 

Here you need to build 2 separate microservices, Order Service and Product Service.

While processing order you will need to fetch details of the products and also manipulate the inventory, so to do that please make use of gRPC apis exposed by Product Service.

Please do not consider this as just an assignment to implement gRPC APIs, but also consider this as you all are building a production grade system.

Focus on following things while you work on it:
- Use both REST(to interact with client app) and gRPC APIs(for inter-service communication)
- Do some research and try to follow a standard project layout while implementing the code, which could scale better when you try to add more APIs in future
- do proper logging and error handling
- try using a real database(sql or nosql, your choice), use any filebased DB like sqlite or boltdb

### Problem Statement:

Please find the below two services and the operations that are allowed.

- Product Service: provides information about the product like availability, price, category

- Order service: provides information about the order like orderValue, dispatchDate, orderStatus, prodQuantity

The user should be able to get the product catalogue and using that info should be able to place an order.

Once the order is placed for a particular product, the product catalogue should be updated accordingly. (Max quantity of a particular product that can be ordered is 10) If the order contains 3 premium different products, order value should be discounted by 10%

The Order service should be able to update the orderStatus for a particular order. dispatchDate should be populated only when the orderStatus is 'Dispatched'.

product category values: Premium/Regular/Budget 
order status values: Placed/Dispatched/Completed/Returned/Cancelled

Postman link : https://api.postman.com/collections/13056412-f14618c5-4121-417a-9848-f268ec4321c5?access_key=PMAT-01H83NAHZ6R6HF96CE40VGNGGY

