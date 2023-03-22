def calculateScoreOfShape(shape):
    if shape == 'A' or shape == 'X':
        return 1
    if shape == 'B' or shape == 'Y':
        return 2
    if shape == 'C' or shape == 'Z':
        return 3


def calculateScore(player1, player2):
    player1Score = calculateScoreOfShape(player1)
    player2Score = calculateScoreOfShape(player2)

    if player1Score == player2Score:
        player1Score += 3
        player2Score += 3

    elif player1 == 'A' and player2 == 'Y':
        player1Score += 6

    elif player1 == 'B' and player2 == 'Z':
        player1Score += 6

    elif player1 == 'C' and player2 == 'X':
        player1Score += 6

    elif player1 == 'B' and player2 == 'X':
        player2Score += 6

    elif player1 == 'C' and player2 == 'Y':
        player2Score += 6

    elif player1 == 'A' and player2 == 'Z':
        player2Score += 6

    return player1Score, player2Score


with open('input.txt') as f:
    lines = f.readlines()
print(len(lines))
totalScorePlayer1 = 0
totalScorePlayer2 = 0
for line in lines:
    players = line.rstrip().split(" ")
    player1 = players[0]
    player2 = players[1]
    player1Score, player2Score = calculateScore(player1, player2)
    totalScorePlayer1 += player1Score
    totalScorePlayer2 += player2Score

print(totalScorePlayer1)
print(totalScorePlayer2)
