# Various terms

Regarding email:

- **MX** : MX record or a mail exchange record is a DNS record that routes emails to the specified email servers. Essentially point to a IP address of the mail server's domain.
- **SPF** : Sender policy framework is a email authentication standard that helps protect the receiving users from spam/spoof/phishing attempts. Adding a SPF to a DNS allows users of that domain to send mail from that domain. Sort of like a permission to the users of that domain to send email.
- **DKIM** : Domain Keys Identified Maill is a authentication method that uses digital keys to indicate that the mail has been verified by the domain.
- **DMARC** : Domain-based Message Authentication, Reporting & Conformance is the sort-of an opposite of an SPF. It protects users from emails. It handles emails either authenticated using a SPF or DKIM or not authenticated. 
- **TXT**: TXT record stores information about other domains. Intented for Domain Admins.

Basically:

- SPF uses DNS to publish the domains, subdomains and mail servers from which authorized email can be sent. 
- DKIM uses DNS to advertise the public keys that can be used to authenticate email messages as having legitimately originated from the domain.
- DMARC uses DNS to advertise the policies that should be applied to email that fails to authenticate with SPF, DKIM or both.

