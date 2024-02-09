## Description
DNS server that having upstream resolver via TOR network. Only sole purpose to get IP from DNS. Packed with ads blocking and DNS query caching.

## Program

### Resolving flow

Here is our resolving flow

```mermaid
flowchart TD
    tor(TOR DNS)
    local(local & cache)
    fallback(fallback to DoT)
    user(End User)
    ah(Access Host)

    user --> ah
    ah --> local
    local -- query not found --> tor
    local -- query found --> return
    tor -- query failed --> fallback
    tor -- query success --> return
    fallback --> return
```

