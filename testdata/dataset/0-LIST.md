# Secret Dataset — File Index

157 fixture files covering secret types across 50+ industries and platforms.
All credentials are **fake** and for testing purposes only.

| File | Secret Types |
|------|-------------|
| `.npmrc` | npm token (`npm_…`), GitHub Packages token |
| `ab_testing_personalization.env` | Optimizely SDK key, VWO, AB Tasty, Dynamic Yield, Kameleoon, Convert Experiments |
| `academic_research_apis.env` | NASA API key, ORCID client secret (`APP-`), Semantic Scholar, NCBI, Scopus/Elsevier, Springer, Census Bureau, data.gov |
| `accounting.env` | QuickBooks refresh token, Xero client secret, FreshBooks access token, Stripe Connect + webhook secret (`whsec_`) |
| `adtech_programmatic.env` | Google Ad Manager refresh token, DV360 service account, The Trade Desk, Criteo, Xandr, Amazon DSP |
| `affiliate_marketing.env` | Impact account SID, ShareASale token, CJ Affiliate, Rakuten Advertising, Awin, PartnerStack, Rewardful |
| `age_ansible.txt` | age encryption key (`AGE-SECRET-KEY-1…`), Ansible Vault encrypted blob, AWX password |
| `agriculture_iot.env` | John Deere Operations Center, Climate Corp, Trimble Agriculture, Arable, Sentera, NREL, FarmLogs |
| `ai_ml_apis.env` | Hugging Face (`hf_`), Replicate (`r8_`), Cohere, Mistral, Groq (`gsk_`), Together AI, Weights & Biases, LangSmith (`ls__`) |
| `alerting_oncall.env` | OpsGenie API/integration key, VictorOps, PagerDuty events v2 key, xMatters, Grafana OnCall token |
| `android_gradle.properties` | Android keystore passwords, Google Play service account JSON |
| `apple_ios.env` | App Store Connect EC private key, Apple team ID, APN cert password |
| `ar_vr_xr.env` | Niantic Lightship, 8th Wall, Snap AR, Meta Spatial app secret, Vuforia server/client keys, Wikitude SDK key |
| `atlassian_okta.env` | Jira/Confluence API tokens (`ATATTx…`), Okta token, PagerDuty key |
| `ats_recruiting.env` | Greenhouse harvest key + webhook secret, Lever, Workable, iCIMS, Ashby, Rippling |
| `auth_identity_providers.env` | Auth0 client secret + management token, Keycloak admin password, FusionAuth, Clerk (`sk_live_`), WorkOS |
| `automotive_connected.env` | Tesla refresh/access token (JWT), Ford Connected, BMW Connected Drive, GM OnStar, Smartcar |
| `aviation_maritime.env` | FlightAware, AviationStack, OpenSky, FlightRadar24, MarineTraffic, VesselTracker, ADS-B Exchange |
| `aws_credentials` | AWS access key (`AKIA…`) + secret access key |
| `aws_sts_session.env` | AWS STS temporary credentials (`ASIA…` + session token), S3 presigned URL with embedded credentials |
| `azure_connection.config` | Azure storage connection string, client secret, tenant ID |
| `azure_sas_cdn.env` | Azure Blob SAS URL (`sig=`), standalone SAS token, Fastly API token, Akamai EdgeGrid credentials |
| `bioinformatics.env` | NCBI/Entrez, Ensembl REST, UniProt, Benchling client secret, Illumina BaseSpace token, DNAnexus auth token |
| `blockchain_defi.env` | Etherscan/BSCScan/PolygonScan API keys, Moralis, The Graph, QuickNode endpoints, Chainlink node password, OpenSea |
| `brand_monitoring.env` | Mention, Brand24, Brandwatch, Talkwalker, Meltwater, Cision, Sprinklr |
| `browser_testing.env` | BrowserStack, Sauce Labs, LambdaTest, Playwright Service token, Percy, Chromatic |
| `cad_3d_manufacturing.env` | Autodesk Forge JWT, OnShape access/secret key, SolidWorks, PTC Windchill password, ANSYS license server, Trimble |
| `cargo_nuget.config` | Cargo registry token, RubyGems API key, NuGet API key |
| `ci_platforms.env` | CircleCI PAT (`ccipat_…`), Travis CI, Jenkins API token, Azure DevOps PAT, TeamCity, Bitbucket app password |
| `climate_environment.env` | OpenAQ, Breezometer, IQAir, PurpleAir read/write keys, Climatiq, Patch carbon offsets, Watershed |
| `cloudflare_newrelic.env` | Cloudflare API token/key, New Relic license + insights key, Sentry auth token (`sntrys_`) |
| `cms_headless.env` | Contentful delivery/management tokens (`CFPAT-`), Sanity auth token, Strapi JWT secrets, Ghost Admin API key |
| `corporate_lms.env` | Docebo client secret, Cornerstone OnDemand, Absorb LMS private/public key, TalentLMS, Litmos, Bridge LMS, 360Learning |
| `crm_sales.env` | Zoho CRM refresh token (`1000.`), Pipedrive, Close.io (`api_`), Copper + user email pair, Freshsales, Affinity |
| `crypto_custody.env` | Fireblocks API key + RSA private key, BitGo access token (`v2x_`), Copper.co, Anchorage, Gemini, Coinbase Prime signing key |
| `crypto_exchanges.env` | Binance API + secret key, Coinbase CDP EC private key, Kraken private key (base64), Bybit secret |
| `crypto_web3.env` | Ethereum private key (64-char hex), BIP-39 mnemonic (12 words), Infura/Alchemy, Pinata IPFS key |
| `customer_support.env` | Zendesk API + OAuth token, Freshdesk, Help Scout, ServiceNow client secret + admin password, Front, Kustomer |
| `data_enrichment.env` | Clearbit (`sk_`), ZoomInfo, Apollo.io, Lusha, Hunter.io, FullContact, PeopleDataLabs, Snov.io |
| `data_stack.env` | Databricks PAT (`dapi_`), dbt Cloud, Airbyte, Fivetran, Segment write key, PostHog (`phc_`), RudderStack |
| `database.yml` | Hardcoded PostgreSQL password in YAML config |
| `design_tools.env` | Figma personal token (`figd_`) + webhook passcode, Miro, InVision, Zeplin, Sketch Cloud, Adobe Creative Cloud |
| `digital_asset_management.env` | Bynder permanent token, Widen access token, Canto app secret, Brandfolder, Acquia DAM, Nuxeo auth token |
| `discord_telegram.env` | Discord bot token, Discord webhook URL, Telegram bot token |
| `django_settings.py` | Django `SECRET_KEY`, hardcoded DB password, SMTP password |
| `dns_domain_registrars.env` | GoDaddy API key/secret, Namecheap, Gandi, OVH app key + consumer key, DNSimple token |
| `docker_config.json` | Docker registry auth (base64-encoded credentials) |
| `ecommerce_platforms.env` | WooCommerce consumer key (`ck_`/`cs_`), BigCommerce access token, Magento consumer key, PrestaShop, OpenCart |
| `education_lms.env` | Canvas LMS access token, Moodle token + admin password, Blackboard, Coursera, Udemy, Duolingo JWT |
| `electronics_components.env` | Digi-Key client secret, Mouser, Octopart, Arrow Electronics, JLCPCB/PCBWay, Altium license key, Cadence license server |
| `email_providers.env` | Mailgun (`key-`), Postmark server/account tokens, Brevo (`xkeysib-`), Klaviyo (`pk_live_`), Resend (`re_`), Loops |
| `energy_utilities.env` | EIA API key, Enphase client secret, SolarEdge, Tesla Powerwall credentials, Octopus Energy (`sk_live_`), NREL |
| `error_tracking.env` | Bugsnag notify key, Rollbar post server item, Raygun, Airbrake, Honeybadger, TrackJS, LogRocket, AppSignal, Embrace |
| `esign_documents.env` | DocuSign integration key + RSA private key, HelloSign, PandaDoc, Adobe Sign (`CBJCHB…`), Lob (`live_`) |
| `expense_finance_ops.env` | Expensify partner secret, SAP Concur, Brex, Ramp, Bill.com API + session ID, Tipalti |
| `feature_flags.env` | LaunchDarkly SDK key (`sdk-`), Split.io, Flagsmith server key (`ser.`), Unleash, GrowthBook |
| `field_service.env` | Jobber access token, Housecall Pro, ServiceMax, FieldAware, Workiz, ServiceTitan, Routific |
| `financial_apis.env` | Plaid access token, Adyen API key + HMAC key, Checkout.com (`sk_`/`pk_`), Klarna basic auth |
| `financial_market_data.env` | Bloomberg API key, Refinitiv/LSEG credentials, IEX Cloud (`sk_`), Quandl, Polygon.io, Tiingo |
| `firebase.js` | Firebase web API key (`AIzaSy…`) + app config |
| `fitness_wellness.env` | Mindbody API key + staff token, ClassPass, Wodify, Zen Planner, PushPress, Strava client secret + refresh token |
| `food_delivery.env` | Yelp Fusion, OpenTable, DoorDash developer ID, Uber Eats, Toast POS restaurant GUID |
| `fraud_detection.env` | Sift, Kount merchant key, Signifyd, Forter site + secret key, MaxMind license key, ThreatMetrix, Neustar |
| `game_platforms.env` | Steam Web API key, Epic Games client secret, PlayFab secret, Unity Cloud Build key, Apple Game Center EC key |
| `gcp_service_account.json` | GCP service account key (full JSON format) |
| `geospatial_maps.env` | HERE Maps API key + app code, TomTom, Esri ArcGIS (`AAPTxy8B…`), What3Words, OpenCage, IPInfo token |
| `github_actions_workflow.yml` | Hardcoded secrets inside a CI workflow `env:` block |
| `github_app.pem` | GitHub App private key PEM + app ID, installation ID, client secret, webhook secret |
| `github_token.env` | GitHub Personal Access Token (`ghp_…`) |
| `gitlab_tokens.env` | GitLab PAT, deploy token, runner token, CI job token (`glpat-`, `gldt-`, `glrt-`) |
| `google_developer_apis.env` | Google Maps/YouTube API keys (`AIzaSy…`), Twitch, Google OAuth client secret (`GOCSPX-`) |
| `graph_specialty_databases.env` | Neo4j Aura URI + password + client secret, CockroachDB connection string, SingleStore, TiDB Cloud, FaunaDB, Dgraph |
| `healthcare_fhir.env` | Epic FHIR private key + client ID, Cerner client secret, Azure Health API, Veeva security token, HL7 MLLP |
| `heroku_digitalocean.env` | Heroku UUID key, DigitalOcean token (`dop_v1_`), Shopify (`shpat_`), Mailchimp key |
| `hospitality_pms.env` | Opera PMS credentials, Mews access + client token, Cloudbeds client secret, Apaleo, Protel, HAPI Hotel |
| `hr_payroll.env` | BambooHR API key, Gusto client secret + token, Workday refresh token, ADP client secret |
| `identity_verification.env` | Onfido (`api_live.`), Jumio, Persona (`persona_sandbox_`), Socure, Checkr, Stripe Identity, Twilio Verify, Authy |
| `image_generation_ai.env` | Stability AI (`sk-`), Midjourney, Remove.bg, DeepAI, Leonardo.ai, Clipdrop, Fal.ai, Ideogram, Photoroom |
| `insurtech.env` | Lemonade, Root Insurance, Socotra host + client secret, Majesco, EZLynx, Applied Epic, Guidewire, Duck Creek |
| `iot_mqtt.env` | MQTT broker credentials, AWS IoT Core cert + private key, Particle.io access token, Twilio Super SIM key |
| `jwt_example.txt` | Raw JWT token (Bearer) + signing secret |
| `kubeconfig.yaml` | Kubernetes bearer token + base64-encoded TLS certs/keys |
| `kubernetes_secret.yaml` | K8s `Secret` manifest with base64-encoded DB/API/JWT/SMTP values + `dockerconfigjson` |
| `laravel.env` | Laravel `APP_KEY` (`base64:…`), DB/mail/Pusher passwords |
| `ldap_active_directory.env` | LDAP bind DN + password, Active Directory service account, Kerberos keytab (base64) |
| `legacy_jdbc_oracle.properties` | Oracle JDBC + TNS password, MSSQL connection string, DB2 JDBC URL, SAP HANA credentials |
| `legal_compliance.env` | LexisNexis client secret, Westlaw, PACER username + client code, CourtListener token, Ironclad, Evisort |
| `legal_practice.env` | Clio client secret + access token, MyCase, PracticePanther, LawPay, Smokeball, FileVine, Litify, Actionstep |
| `localization_i18n.env` | Crowdin personal token, Lokalise, Phrase, Transifex (`1/`), Weblate, POEditor, Smartling user secret |
| `log_management.env` | Papertrail token, Loggly customer + API token, Mezmo/LogDNA, Coralogix private key, Logz.io, Sumo Logic, Grafana Loki |
| `mapbox_algolia.env` | Mapbox tokens (`pk.eyJ…`/`sk.eyJ…`), Algolia admin key, Amplitude, Mixpanel |
| `maven_settings.xml` | Nexus deploy password + GitHub Packages PAT in Maven config |
| `media_streaming.env` | Mux token, Cloudinary API secret + URL scheme, Agora app certificate, LiveKit API secret, ImageKit private key |
| `messaging_brokers.env` | Kafka SASL/SSL passwords, Confluent Cloud API key, NATS NKey seed + JWT, InfluxDB token |
| `modern_databases.env` | Supabase anon + service role JWT keys, PlanetScale (`pscale_pw_`), Neon Postgres URL, Turso libSQL token |
| `mongodb_redis.env` | MongoDB Atlas URI, Redis URL, Elasticsearch API key, RabbitMQ AMQP URL |
| `multicloud_providers.env` | Oracle Cloud OCI fingerprint + key, IBM Cloud API key, Alibaba Cloud (`LTAI5t…`), Hetzner Cloud, Scaleway |
| `music_audio.env` | Apple MusicKit EC private key + key ID, SoundCloud OAuth token (`2-`), Last.fm shared secret, Discogs, Bandcamp |
| `netrc` | `.netrc` machine credentials for multiple hosts |
| `network_infrastructure.env` | Cisco IOS enable password + DNAC token, Palo Alto API key, FortiGate, F5 BIG-IP, Juniper NETCONF |
| `object_storage_s3compat.env` | Backblaze B2 application key, Wasabi, Cloudflare R2 access key, MinIO root credentials, Storj access grant |
| `observability.env` | Grafana service account token (`glsa_`), Dynatrace API/PaaS tokens (`dt0c01.`), Splunk HEC token, Datadog app key |
| `open_banking.env` | TrueLayer client secret, Tink, Salt Edge app secret, Belvo secret password, Nordigen secret key, Yapily |
| `openai_config.py` | OpenAI API key (`sk-proj-…`) + Anthropic API key (`sk-ant-…`) |
| `ota_firmware.env` | Mender.io access token, Balena, AWS IoT Device Management, Azure IoT Hub connection string, secure boot signing key |
| `password_managers.env` | 1Password service account token (`ops_`), Bitwarden client secret, Dashlane, LastPass MFA secret, Vault token (`hvs.`) |
| `payment_gateways.env` | Braintree merchant/public/private keys, PayPal client secret, Square access + app secret (`EAAA`, `sq0csp-`) |
| `performance_testing.env` | k6 Cloud token, BlazeMeter key/secret, Gatling Enterprise token, Artillery Cloud, LoadRunner license server, Flood.io, NeoLoad |
| `pgp_private.asc` | PGP private key block |
| `pki_certificates.env` | DigiCert API key, Sectigo cert password, GlobalSign API secret, ACME EAB key ID + HMAC, PKCS#12 passphrase, Venafi |
| `podcast_streaming.env` | Transistor, Buzzsprout + podcast ID, Simplecast, Podbean, Castos, Captivate + user ID, RSS.com, Anchor |
| `private_key.pem` | RSA private key PEM block |
| `project_management.env` | Asana access token, Monday.com JWT-style key, ClickUp (`pk_`), Basecamp, Wrike, Trello API key + token |
| `pyproject_ci.toml` | PyPI token (`pypi-Ag…`), SonarQube token (`sqp_`), Codecov token, Pulumi access token (`pul-`) |
| `quantum_computing.env` | IBM Quantum token, IonQ API key, D-Wave Leap token, Azure Quantum client secret, AWS Braket keys, Rigetti |
| `real_estate.env` | Zillow, Redfin, Realtor.com, MLS/RETS credentials, RESO bearer token, ATTOM Data, CoreLogic |
| `revenue_intelligence.env` | Gong access key + secret, Chorus AI, Outreach client secret, Salesloft, Clari, People.ai, Wingman |
| `robotics_industrial.env` | ROS master URI + security keystore, ABB/Universal Robots/Fanuc/Kuka controller passwords, Siemens MindSphere, PTC Kepware, Ignition Gateway |
| `salesforce_hubspot.env` | Salesforce OAuth client secret + password+token, HubSpot private app token (`pat-na1-`), Intercom |
| `satellite_space.env` | Planet Labs, Maxar client secret, Copernicus (`sh-`) client secret, Spire, AWS Ground Station role ARN, NOAA CDO token |
| `security_edr.env` | CrowdStrike client secret, SentinelOne management URL + token, Tenable access/secret keys, Rapid7, Qualys, Carbon Black |
| `security_scanners.env` | Snyk token, Aqua Security, Prisma Cloud, JFrog Artifactory token (`cmVmdGtu`), Nexus, Harbor |
| `sftp_ftp_config.json` | SFTP config (VSCode extension format) — password + private key passphrase |
| `shipping_logistics.env` | Shippo (`shippo_live_`), EasyPost (`EZAK`), ShipStation, FedEx, UPS, DHL API keys |
| `slack_webhook.py` | Slack webhook URL + bot token (`xoxb-`) |
| `smart_home_iot.env` | Home Assistant long-lived token (JWT), Philips Hue bridge username + client key, Nest project ID, Samsung SmartThings, Tuya |
| `snowflake_redshift.env` | Snowflake private key + password, Redshift DSN, inline BigQuery credentials |
| `social_media_management.env` | Hootsuite, Buffer, Sprout Social, Later, Publer, SocialBee access tokens |
| `social_niche_platforms.env` | Mastodon access token + client key/secret, Bluesky app password (`xxxx-xxxx-xxxx-xxxx`), Pinterest, LinkedIn (`AQ`), Threads |
| `sports_events.env` | Sportradar, Ticketmaster key/secret, Eventbrite private token, SeatGeek, ESPN, Stats Perform, Football-Data.org |
| `spring_application.properties` | Spring datasource/mail passwords, AWS keys in Java config, JWT signing key |
| `ssh_config` | Hardcoded SSH password in config |
| `stock_media.env` | Shutterstock token (`v2/`), Getty Images client secret, Unsplash access/secret key, Pexels, Adobe Stock, Pixabay |
| `stripe_key.js` | Stripe live secret key (`sk_live_…`) |
| `subscription_billing.env` | Chargebee (`live_`), Recurly + public key, Paddle vendor auth code + public key, Zuora, Maxio, Stripe billing webhook |
| `supply_chain_procurement.env` | SAP Ariba API key + realm, Coupa client secret + instance URL, Oracle Procurement, Jaggaer, Ivalua, GEP, Tradogram |
| `surveys_forms.env` | Typeform personal token (`tfp_`), SurveyMonkey, JotForm, Qualtrics, Google Forms (`AIzaSy`), Tally, Paperform |
| `tailscale_vpn.env` | Tailscale auth key (`tskey-auth-`), ZeroTier token, OpenVPN static key, Cloudflare Tunnel token |
| `tax_compliance.env` | Avalara account ID + license key, TaxJar token, Vertex client secret, Sovos, Thomson Reuters ONESOURCE, Vatstack |
| `telco_sms.env` | Vonage API secret + app private key, MessageBird, Sinch app key/secret, Bandwidth, Plivo auth ID/token |
| `terraform.tfvars` | AWS keys, Datadog API/app keys, HashiCorp Vault token (`hvs.`) |
| `test` | HTTP proxy URL with embedded credentials |
| `time_tracking.env` | Toggl, Harvest account ID + access token, Clockify, Time Doctor, Hubstaff, RescueTime, Everhour |
| `translation_speech.env` | DeepL (`fx_`), Azure Cognitive/Speech/Translator keys, Google Speech, Deepgram, AssemblyAI, ElevenLabs |
| `travel_booking.env` | Amadeus client secret, Sabre (`V1:`), Booking.com credentials, Expedia, Airbnb, Skyscanner |
| `twilio_sendgrid.env` | Twilio SID/token + SendGrid API key (`SG.`) |
| `twitter_spotify.env` | Twitter/X full OAuth credential set (bearer + access token), Spotify client secret, Reddit credentials |
| `uptime_status.env` | UptimeRobot (`ur2114919-`), Better Uptime, Pingdom, Statuspage API key + page ID, Checkly (`cu_`), Hyperping |
| `vault_consul.hcl` | Consul ACL token, HashiCorp Vault AppRole `role_id`/`secret_id`, Nomad ACL token |
| `vector_databases.env` | Pinecone (`pcsk_`), Weaviate, Qdrant, Chroma, Typesense admin key, Meilisearch master key, Zilliz |
| `vercel_netlify.env` | Vercel token, Netlify auth token, Fly.io token (`FlyV1 …`), Railway token |
| `video_platforms.env` | Vimeo access token, Brightcove client secret + account ID, JW Player secret/property key, Wistia, Kaltura admin secret |
| `voip_telephony.env` | Asterisk AMI secret, FreeSWITCH ESL password, SIP trunk credentials, RingCentral JWT, Dialpad |
| `vault_consul.hcl` | Consul ACL token, HashiCorp Vault AppRole `role_id`/`secret_id`, Nomad ACL token |
| `weather_news_apis.env` | OpenWeatherMap, WeatherStack, Tomorrow.io, NewsAPI, Guardian, NYT, Alpha Vantage, Finnhub |
| `wireguard.conf` | WireGuard `PrivateKey` + `PresharedKey` |
| `wordpress_wp-config.php` | WordPress DB password, auth/nonce salts, AWS S3 keys in PHP `define()` |
| `zoom_linear.env` | Zoom API key/secret/webhook, Linear API key (`lin_api_`), Notion token (`secret_`), Dropbox token (`sl.`) |
