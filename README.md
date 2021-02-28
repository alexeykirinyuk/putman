# putman

# db structure
- collection
    - id (guid - pk)
- folder
    - id (guid - pk)
    - name (str)
    - collection id (guid - fk on collection)
    - parent id (nullable guid - fk on folder)
- request
    - id (guid - pk)
    - name (str)
    - method (method)
    - content (str nullable)
    - body (str)
    - folder id (nullable guid - fk)
- header
    - id (guid - pk)
    - request id (guid - fk on request)
    - name (str)
    - value (str)

# putman v1
- list of collections
- collection describe (full view)
- add/remove/rename collection
- add/remove/rename folder
- add/remove/rename request
- edit request (touch mode)