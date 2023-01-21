import subprocess
import tempfile

from revChatGPT.ChatGPT import Chatbot
from loguru import logger
import fire


class Bot(object):
    def __init__(self, config: dict):
        chatbot = Chatbot(config)
        chatbot.reset_chat()
        chatbot.refresh_session()
        self._bot = chatbot
        self._conversation_id = None
        logger.info("chatbot ready")

    def get_resp(self, req: str) -> str:
        resp = self._bot.ask(req, conversation_id=self._conversation_id)
        # update
        self._conversation_id = resp["conversation_id"]
        return resp["message"]

    def reset(self):
        self._conversation_id = None


def new_bot(session_token: str) -> Bot:
    config = {
        "session_token": session_token,
    }

    return Bot(config)


class Topic(object):
    def __init__(self):
        self.title: str = ""
        self.content: str = ""

    def to_md(self) -> str:
        return f"""
# {self.title}

{self.content}

"""


def main(
        token: str = "token.txt",
        input_file: str = "data.txt",
        output_file: str = "output.html",
        marp_exe: str = "./marp"):
    with open(token, encoding="utf-8") as f:
        session_token = f.read()

    with open(input_file, encoding="utf-8") as f:
        question_list = f.readlines()

    topic_list = []

    bot = new_bot(session_token)
    for each in question_list:
        resp = bot.get_resp(each)
        topic = Topic()
        topic.title = each
        topic.content = resp
        topic_list.append(topic)

    # gen
    parts = [each.to_md() for each in topic_list]
    final = "\n---\n".join(parts)
    logger.info(f"gen finished, turn it to html: {final}")

    with tempfile.NamedTemporaryFile(mode="w", encoding="utf-8") as tmp:
        tmp.write(final)
        tmp.flush()

        try:
            subprocess.check_call([marp_exe, tmp.name, "-o", output_file])
        except FileNotFoundError:
            logger.error("call marp failed. generate raw markdown.")
            with open(output_file, "w+", encoding="utf-8") as f:
                f.write(final)

    logger.info(f"result saved to {output_file}")


def cmd():
    fire.Fire(main)


if __name__ == '__main__':
    cmd()
