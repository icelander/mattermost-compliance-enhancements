# Mattermost Compliance Enhancements

# About

This plugin adds functionality to make compliance exports easier to generate and track. At the moment it only adds the `editedBy` property to a post if the editor is different from the user who posted the message, though more functionality may be added in the future.

## Installation

1. Read the [Changelog](#changelog)
2. Download the latest release from [Releases](./releases)
3. Upload the `tar.gz` file to your server
4. Enable the plugin

## Changelog

### [0.1.0] - 2021-11-23
#### Added
- Records the ID of the user who edited a post in the properties field under `editedBy`