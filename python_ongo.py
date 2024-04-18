from pymongo import MongoClient

# Connect to MongoDB
client = MongoClient('mongodb://199.241.138.96:27017/')
db = client['gotest']  # Replace 'mydatabase' with the name of your database
collection = db['ddusers']  # Replace 'mycollection' with the name of your collection

# Insert a document into the collection
document = {'name': 'John', 'age': 30}
result = collection.insert_one(document)
print("Inserted document ID:", result.inserted_id)

# Find documents in the collection
results = collection.find()
print("Found documents:")
for result in results:
    print(result)

# # Update documents in the collection
# update_query = {'name': 'John'}
# update_data = {'$set': {'age': 35}}
# result = collection.update_one(update_query, update_data)
# print("Modified document count:", result.modified_count)

# # Fi
