# golarm

The simplest possible tool I could come up with to create little reminders
leveraging the systems native notifications.

![Notification](/notification.png)

## Building from source

```bash
go build .
cp golarm /usr/local/bin
```

## Creating a notification

```shell
# A notification which triggers every day at 09:30h
golarm -cron="30 09 * * 1-5" -message="Daily meeting"

# A one-off appointmet the 15th of February at 17:35h
golarm -cron="35 17 15 02 *" -message="Appointment at dentist"
```

## crontab

golarm simply leverages [`crontab`][0] for all heavy-lifting, so to clear all
notifications you can just run `crontab -r`.

Likewise, to check what notifications you have enabled, you can run
`crontab -l`.

[0]: http://man7.org/linux/man-pages/man5/crontab.5.html
