from enum import Enum

class Choice(Enum):
    ROCK = 1
    PAPER = 2
    SCISSOR = 3
    def beats(self):
        if self == Choice.ROCK:
            return Choice.SCISSOR
        if self == Choice.PAPER:
            return Choice.ROCK
        if self == Choice.SCISSOR:
            return Choice.PAPER
    def beaten_by(self):
        for option in Choice:
            if option.beats() == self:
                return option


def score_in_ideal_game(input):
    score = 0
    for line in iter(input.splitlines()):
        split_line = line.split()
        opponent_choice = get_choice(split_line[0])
        my_choice = get_choice(split_line[1], options=("X", "Y", "Z"))
        score += score_for_round(opponent_choice, my_choice)
    return score

def get_choice(input, options=("A", "B", "C")):
    if input == options[0]:
        return Choice.ROCK
    elif input == options[1]:
        return Choice.PAPER
    elif input == options[2]:
        return Choice.SCISSOR
    else:
        raise ValueError("Invalid choice: " + input)

def score_in_game_with_target_result(input):
    score = 0
    print(len(input.splitlines()))
    for line in iter(input.splitlines()):
        split_line = line.split()
        opponent_choice = get_choice(split_line[0])
        if split_line[1] == "Y":
            my_choice = opponent_choice
        elif split_line[1] == "X":
            my_choice = opponent_choice.beats()
        else:
            my_choice = opponent_choice.beaten_by()
        score += score_for_round(opponent_choice, my_choice)
    return score

def score_for_round(opponent_choice, my_choice):
    score = my_choice.value

    if opponent_choice == my_choice:
        score +=3

    if my_choice.beats() == opponent_choice:
            score += 6

    return score

if __name__ == '__main__':
    with open('input.txt', 'r') as file:
        input = file.read()
    print(score_in_ideal_game(input))
    print(score_in_game_with_target_result(input))
