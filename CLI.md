# Protomicro

## Generate new proto

protomicro new app
protomicro gen --module gopackage__
protomicro gen --module order/item/return
protomicro gen --module category
protomicro gen --module product
protomicro gen --module payment

## Generate pb, dto, injector, service, model, test from proto file

protomicro gen --module gopackage__

```
pluralName := util.ToPlural(name)
```