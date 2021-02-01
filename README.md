# Goma

Chibi Goma Project Template for [protomicro](https://github.com/gomaglev/protomicro).

## Variables

### proto file

``` Javascript
//pbpackage__                   order.item.message
//gopackage__                   order/item/message
proto.gopackage_name__        order.item.message
proto/gopackage_name__        proto/order/item/message
gofile_name__                 order_item_message
gopackage_name__              message
version__                     v1
pb_name__                     message
PbName__                      Message
pb_name_plural__              messages
PbNamePlural__                Messages
gateway_uri__                 v1/orders/{order_id}/items/{item_id}/messages
MessageTypeName__             OrderItemMessage
MessageTypeNamePlural__       OrderItemMessages
table_name__                  order_item_message
//ApiProtoParentIDs__         
    /*
        string order_id = 3;
        string item_id = 4; 
    */

//TypeProtoParentIDs__        
    /* 
    // @inject_tag: faker:"uuid_hyphenated"
    string order_id = 3;
    string item_id = 4; 
    */

//DtoProtoParentIDs__        
    /* 
    OrderId    string
    ItemId     string
    */

```

### app

``` Javascript
app_module__                  goma / github.com/gomaglev/goma
app_name__                    goma
AppNameTitleCase__            Goma
```
