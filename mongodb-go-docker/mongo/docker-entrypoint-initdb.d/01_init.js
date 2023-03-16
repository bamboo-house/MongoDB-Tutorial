var user = {
  user: "mongo",
  pwd: "mongo",
  roles: [
    {
      role: "dbOwner",
      db: "sample",
    },
  ],
};

db.createUser(user);
db.createCollection("sample_collection");
