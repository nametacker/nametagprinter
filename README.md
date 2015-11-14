# Nametag printer

[![Travis](https://travis-ci.org/coderbyheart/nametagprinter.svg?branch=master)](https://travis-ci.org/coderbyheart/nametagprinter/)

This is a webservice for printing nametags on events.

## Testing

Run the tests

    go test ./...

## Setting the paperformat for label printers

### Brother QL-720NW

Use the `brpapertoollpr_ql720nw` program.

    # Add a new paper format called BCRM which is 80x50mm
    sudo brpapertoollpr_ql720nw -P Brother_QL-720NW -n 50x80 -w 50 -h 80
    # Ensure that it has been added
    lpoptions -d Brother_QL-720NW -l
    # It's added to the end of the PageSize/Media Size, like BrL05500328296A
    # and can now selected in the Default Options for the printer in cups
