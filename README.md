# boggart

Highly customizable low-interaction experimental honeypot that mimics specific hosts.

<a href="https://github.com/edoardottt/boggart/actions">
 <img src="https://github.com/edoardottt/boggart/actions/workflows/go.yml/badge.svg" alt="workflows" />
 <img src="https://goreportcard.com/badge/github.com/edoardottt/boggart" alt="goreportcard" />
</a>
<br><br>

*'So, the first question we must ask ourselves is, what is a Boggart?'*  
Hermione put up her hand.  
*'It's a shape-shifter'*, she said. *'It can take the shape of whatever it think will frighten us most'*.  
*'Couldn't have put it better myself'*, said Professor Lupin, and Hermione glowed.  
*'So the Boggart sitting in the darkness within has not yet assumed a form. He does not yet know what will frighten the person on the other side of the door. Nobody knows what a Boggart looks like when he is alone, but when I let him out, he will immediately become whatever each of us most fears'*.

<p style="text-align:right;">Harry Potter and the Prisoner of Azkaban</p>

Installation 📡
-------

*Disclaimer*: This is an ongoing and experimental project: there are features not yet available and features not (fully) tested. It is designed for home labs / home environments, not for professional or industrial purposes. Deploy in your network at your own risk.

[Docker](https://docs.docker.com/get-docker/) and [Docker compose](https://docs.docker.com/compose/install/) are needed.

```console
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

Read the [docs](https://github.com/edoardottt/boggart/tree/main/docs) to understand how it works and how to configure your honeypot.

Changelog 📌
-------

Detailed changes for each release are documented in the [release notes](https://github.com/edoardottt/boggart/releases).

Contributing 🤝
------

If you want to contribute to this project, you can start opening an [issue](https://github.com/edoardottt/boggart/issues).

Before opening a pull request, download [golangci-lint](https://golangci-lint.run/usage/install/) and run

```bash
golangci-lint run
```

If there aren't errors, go ahead :)

License 📝
-------

This repository is under [GNU General Public License v3.0](https://github.com/edoardottt/boggart/blob/main/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.

Created with [gonesis](https://github.com/edoardottt/gonesis)❤️
