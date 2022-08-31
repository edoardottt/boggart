# boggart
Highly customizable low-interaction experimental honeypot that mimics specific hosts, behaviors, and CVEs.

*Disclaimer*: This is an ongoing and experimental project: this means there are features not yet available and features not (fully) tested. It is designed for home labs / home environments, not for professional or industrial purposes. Deploy in your network at your own risk.

Installation 📡
-------

[Docker](https://docs.docker.com/get-docker/) and [Docker compose](https://docs.docker.com/compose/install/) are needed.

```
git clone https://github.com/edoardottt/boggart
```

Usage 💻
-------

- Edit the configuration file `config.yaml` setting up the machine you want to create
- Execute `make up` (inside the boggart folder)

Now you have three open ports on your local machine:
  - [localhost:8092](http://localhost:8092/) - This is the actual honeypot
  - [localhost:8093](http://localhost:8093/) - This is the dashboard (do not expose this !)
  - [localhost:8094](http://localhost:8094/) - This is the API service (do not expose this !)
  
You must expose on the public Internet only the service hosted on port 8092.

Changelog 📌
-------
Detailed changes for each release are documented in the [release notes](https://github.com/edoardottt/boggart/releases).

Contributing 🤝
------
If you want to contribute to this project, you can start opening an [issue](https://github.com/edoardottt/boggart/issues).

License 📝
-------
This repository is under [GNU General Public License v3.0](https://github.com/edoardottt/boggart/blob/main/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.

Created with [gonesis](https://github.com/edoardottt/gonesis)❤️
