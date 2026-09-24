# Port taxonomy

Reference catalog for Kestrel's enumeration module (`internal/enum`), organized by category rather than a flat list — matching the Discovery → Enumeration → Detection → Validation separation that shapes the rest of the platform.

**A port number never confirms a service.** It is only an indication; the actual service should be confirmed via banner, handshake, or fingerprinting. `internal/enum` does passive banner grabbing today; active fingerprinting is a future enhancement.

## Currently scanned (TCP connect scan, 53 ports)

### Web
| Port | Service |
|---|---|
| 80 | HTTP |
| 443 | HTTPS |
| 3000, 5000 | Common dev server ports (Node.js, Flask, etc.) |
| 8000, 8080 | HTTP / HTTP proxy |
| 8443 | HTTPS alternate |
| 8888 | HTTP alternate (Jupyter and similar) |

### Remote access
| Port | Service |
|---|---|
| 22 | SSH |
| 23 | Telnet |
| 3389 | RDP |
| 5900 | VNC |
| 5985 / 5986 | WinRM (HTTP / HTTPS) |

### File sharing
| Port | Service |
|---|---|
| 20 | FTP-data |
| 21 | FTP |
| 139 | NetBIOS-SSN (SMB over NetBIOS) |
| 445 | SMB |
| 873 | rsync |
| 2049 | NFS |

### VoIP (TCP transport)
| Port | Service |
|---|---|
| 5060 | SIP (TCP) |
| 5061 | SIP-TLS |

> SIP is commonly associated with UDP, but SIP-over-TCP is standard (RFC 3261) and widely deployed — these two ports are TCP-connect scannable, unlike the rest of the VoIP/VPN group below.

### Directory / authentication
| Port | Service |
|---|---|
| 88 | Kerberos |
| 135 | MSRPC |
| 137 / 138 | NetBIOS-NS / NetBIOS-DGM |
| 389 | LDAP |
| 636 | LDAPS |

### Databases
| Port | Service |
|---|---|
| 1433 | MS SQL Server |
| 1521 | Oracle DB |
| 3306 | MySQL / MariaDB |
| 5432 | PostgreSQL |
| 6379 | Redis |
| 9200 / 9300 | Elasticsearch (API / cluster) |
| 27017 | MongoDB |

### Infrastructure & platform services
| Port | Service |
|---|---|
| 53 | DNS |
| 111 | RPCbind |
| 514 | Syslog |
| 902 | VMware/ESXi |
| 1080 | SOCKS proxy |
| 2181 | Apache ZooKeeper |
| 2375 / 2376 | Docker API (plain / TLS) — **2375 without TLS is a critical exposure** if reachable |
| 5601 | Kibana |
| 6443 | Kubernetes API |
| 11211 | Memcached |
| 50000 | SAP and similar enterprise apps |

> SNMP (161/162) is intentionally not listed here as TCP: it's fundamentally a UDP protocol. A TCP connect attempt to 161 will almost always fail even when the SNMP agent is running — it's listed once below, under UDP, where it actually belongs. (An earlier draft of this table listed it in both places, which was a mistake worth flagging rather than quietly copying forward.)

### Mail
| Port | Service |
|---|---|
| 25 | SMTP |
| 110 | POP3 |
| 143 | IMAP |
| 465 | SMTPS |
| 587 | SMTP submission |
| 993 | IMAPS |
| 995 | POP3S |

## Not currently scanned: UDP services

The scanner is TCP-connect only — UDP has no handshake to confirm an open port. Detecting it properly needs either a protocol-specific probe with a real response to interpret (a real DNS query for 53, an SNMP `GetRequest` for 161, a valid NTP packet for 123) or listening for an ICMP "port unreachable" response, which requires a raw socket — on Linux, that means root or `CAP_NET_RAW`, a privilege nothing else in Kestrel currently needs. This is a genuinely different scanning subsystem, not a handful of extra map entries, so it's tracked here as planned work rather than implemented:

| Port | Service | Why it matters | Planned probe complexity |
|---|---|---|---|
| 53 | DNS | Zone transfer, misconfiguration | Low — well-defined query/response |
| 123 | NTP | Infrastructure enumeration | Low — well-defined query/response |
| 161/162 | SNMP / Trap | Device info, weak community strings | Low — well-defined query/response |
| 67/68 | DHCP | Network infrastructure | Medium |
| 137/138 | NetBIOS | Windows enumeration | Medium |
| 1434 | MS SQL Browser | SQL Server discovery | Medium |
| 1900 | SSDP/UPnP | Device discovery | Medium |
| 69 | TFTP | No strong authentication | Medium |
| 500 / 4500 | IKE / IPsec NAT-T | VPN | High |
| 1194 | OpenVPN | VPN | High |

**First candidates for a future `internal/enum/udp` module**: DNS, NTP, and SNMP — all three have a simple, well-documented request/response contract, making them the highest-value, lowest-effort additions when this work is picked up. SIP (5060/5061) is not in this table — see the VoIP section above, since SIP-over-TCP is already covered by the current scanner.

## Planned data model: richer per-port metadata

Today, `internal/enum`'s port list is a simple `map[int]string` (port → service name) — deliberately, to avoid a schema change (`assets.category`, `assets.recommended_tool`, `assets.risk_hint`) not yet justified by a concrete consumer.

Once the Tier 1-3 detector phases (see the main roadmap, phases 08b-08e) are underway, revisit this: a richer catalog of `port → service → protocol → category → recommended enumeration technique → tool → potential risk` would let the Vulnerability Assessment stage prioritize which assets to run which detectors against, instead of running every detector against every asset. Categories such as `web`, `remote-access`, `file-sharing`, `directory-auth`, `database`, `infrastructure`, and `mail` (as used in the tables above) are the natural starting taxonomy for that column.
