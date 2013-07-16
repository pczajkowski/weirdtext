import re
import random


TOKENIZE_RE = re.compile(r'(\w+)', re.U)


def encode_text(text):
    output = TOKENIZE_RE.split(text)

    original_words = []
    encoded_words = []

    for word in output:
        encoded_word, encoded = encoder(word)
        if encoded:
            original_words.append(word)
        encoded_words.append(encoded_word)

    separator = '\n---weird---\n'
    original = sorted(original_words, key=lambda s: s.lower())
    original_words = ' '.join(original)
    encoded_text = ''.join(encoded_words)
    return ''.join([separator, encoded_text, separator, original_words])


def encoder(word):
    if len(word) > 3:
        first_letter = word[0]
        last_letter = word[-1]
        word_center = word[1:-1]
        encoded_center, encoded = shuffle(word_center)

        encoded_word = ''.join([first_letter, encoded_center, last_letter])
    else:
        encoded_word = word
        encoded = False
    return encoded_word, encoded


def shuffle(word):
    encoded = word
    word_list = list(word)
    encoded_list = word_list[:]

    encoded = False
    if len(set(encoded_list)) >= 2:
        while (encoded_list == word_list):
            random.shuffle(word_list)
            encoded = True

    encoded_word = ''.join(word_list)
    return encoded_word, encoded


def decoder():
    pass


if __name__ == '__main__':
    text = '''This is a long looong test sentence,
with some big (biiiiig) words!'''
    print encode_text(text)
