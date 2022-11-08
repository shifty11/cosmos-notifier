import 'package:cosmos_notifier/api/protobuf/dart/admin_service.pbgrpc.dart';
import 'package:cosmos_notifier/common/header_widget.dart';
import 'package:cosmos_notifier/config.dart';
import 'package:cosmos_notifier/f_admin/widget/services/admin_provider.dart';
import 'package:cosmos_notifier/f_home/services/message_provider.dart';
import 'package:cosmos_notifier/f_home/widgets/subwidgets/bottom_navigation_bar_widget.dart';
import 'package:cosmos_notifier/style.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_messages/riverpod_messages.dart';

class AdminPage extends StatelessWidget {
  AdminPage({Key? key, errorCode}) : super(key: key);

  final TextEditingController messageController = TextEditingController();

  final String discordHelp = '''*italics* or _italics_
**bold**
***bold italics***
__underline__

`code block`

```
multiline
```''';
  final String telegramHelp = '''<b>bold</b>
<i>italic</i>
<code>code</code>
<s>strike</s>
<u>underline</u>
<pre language="bash">code</pre>

<a href='https://telegram.org'>Telegram</a>''';

  Widget testButton(BuildContext context, IconData icon, String receiver, Color color, BroadcastMessageRequest_MessageType type) {
    const buttonWith = 170.0;
    return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
      return OutlinedButton.icon(
        onPressed: () => ref.read(adminProvider.notifier).broadcastMessage(messageController.text, type),
        icon: Icon(icon),
        label: Text(receiver),
        style: OutlinedButton.styleFrom(
          minimumSize: const Size(buttonWith, 50),
          primary: color,
          side: BorderSide(color: color),
        ),
      );
    });
  }

  Widget button(BuildContext context, IconData icon, String receiver, Color color, BroadcastMessageRequest_MessageType type) {
    const buttonWith = 170.0;
    return ElevatedButton.icon(
      onPressed: () async => showConfirmSendMessageDialog(context, receiver, type),
      icon: Icon(icon),
      label: Text(receiver),
      style: ElevatedButton.styleFrom(
        minimumSize: const Size(buttonWith, 50),
        primary: color,
        onPrimary: Colors.white,
      ),
    );
  }

  Widget buttons(BuildContext context) {
    var buttons = <Widget>[
      testButton(context, Icons.telegram, "Telegram Test", Styles.telegramColor, BroadcastMessageRequest_MessageType.TELEGRAM_TEST),
      button(context, Icons.telegram, "Telegram", Styles.telegramColor, BroadcastMessageRequest_MessageType.TELEGRAM),
    ];
    if (jwtManager.isDiscordUser) {
      buttons = [
        testButton(context, Icons.discord, "Discord Test", Styles.discordColor, BroadcastMessageRequest_MessageType.DISCORD_TEST),
        button(context, Icons.discord, "Discord", Styles.discordColor, BroadcastMessageRequest_MessageType.DISCORD),
      ];
    }
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: buttons,
    );
  }

  void showConfirmSendMessageDialog(BuildContext context, String receiver, BroadcastMessageRequest_MessageType type) {
    showDialog(
      context: context,
      builder: (_) {
        return Consumer(builder: (BuildContext context, WidgetRef ref, Widget? child) {
          return AlertDialog(
            title: Text('Send message to $receiver'),
            content: SingleChildScrollView(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text("Are you sure you want to send this message to all $receiver users?"),
                ],
              ),
            ),
            actions: [
              TextButton(
                onPressed: () => Navigator.pop(context),
                child: const Text('Cancel'),
              ),
              ElevatedButton(
                style: ElevatedButton.styleFrom(primary: Styles.dangerBgColor, onPrimary: Styles.dangerTextColor),
                onPressed: () {
                  ref.read(adminProvider.notifier).broadcastMessage(messageController.text, type);
                  Navigator.pop(context);
                },
                child: const Text('Send'),
              )
            ],
          );
        });
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [
          MessageOverlayListener(
            provider: messageProvider,
            child: Container(
              width: MediaQuery.of(context).size.width,
              padding: const EdgeInsets.only(top: Styles.topPadding, left: Styles.sidePadding, right: Styles.sidePadding),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const HeaderWidget(),
                  const SizedBox(height: 10),
                  ConstrainedBox(
                    constraints: BoxConstraints(
                      maxHeight: MediaQuery.of(context).size.height - 200,
                    ),
                    child: TextField(
                      controller: messageController,
                      keyboardType: TextInputType.multiline,
                      minLines: 1,
                      maxLines: null,
                      decoration: InputDecoration(
                        border: OutlineInputBorder(),
                        // floatingLabelAlignment: FloatingLabelAlignment.center,
                        alignLabelWithHint: true,
                        labelText: 'Message',
                        hintText: jwtManager.isTelegramUser ? telegramHelp : discordHelp,
                      ),
                    ),
                  ),
                  const SizedBox(height: 10),
                  buttons(context),
                ],
              ),
            ),
          ),
        ],
      ),
      bottomNavigationBar: BottomNavigationBarWidget(jwtManager: jwtManager),
    );
  }
}
