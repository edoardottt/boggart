# Guide

This is a brief description of how boggart works and how you can configure it. If you think something is missing or you are not understanding something just send me a message wherever you want.  

Why is it highly customizable? Because the honeypot configuration file is a YAML file and you can configure how it must work (see TEMPLATE_GUIDE.md to understand how you have to write the template file in order to build your desired honeypot).

When it starts, it exposes three ports:
  - 8092: This is the actual honeypot
  - 8093: This is the dashboard (do not expose this !)
  - 8094: This is the API service (do not expose this !)

You must expose on the public Internet only the service hosted on port 8092.  
To understand how API works see API.md
