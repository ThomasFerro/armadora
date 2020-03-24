# Armadöra

![Game in progress](https://github.com/ThomasFerro/readmes/blob/master/posts/12-ddd-in-action-armadora-the-board-game/game_in_progress.jpg)

Armadora is a game where every **player** will try to get their hands on the most **gold** possible.

![Board](https://github.com/ThomasFerro/readmes/blob/master/posts/12-ddd-in-action-armadora-the-board-game/board.jpg)

The **game board** is a 8x5 grid with two types of **cell**: the **lands** where the players will be able to put their **warriors** (more on that later) and the **gold stacks** with various quantity or gold.

The grid can be divided into **territories** by putting **palisades** between cells. A territory *must be at least four squares wide*.

At the beginning of a game, each player chooses a character (*Orc*, *Goblin*, *Elf* or *Mage*) and hide their warriors behind their screens.

Each player starts with the same army, depending on the number of players:

- For two players, each one get *11 warriors of 1 point*, *2 warriors of 2 points*, *1 warrior of 3 points*, *1 warrior of 4 points* and *1 warrior of 5 points*;
- For three players, each one get *7 warriors of 1 point*, *2 warriors of 2 points*, *1 warrior of 3 points* and *1 warrior of 4 points*;
- For four players, each one get *5 warriors of 1 point*, *1 warrior of 2 points*, *1 warrior of 3 points* and *1 warrior of 4 points*;

In the advanced rules set, the players also have power token based on their race and a reinforcement token.

*Forty gold* are then distributed randomly in eight piles: 1 pile of 3, 2 piles of 4, 2 piles of 5, 2 piles of 6 and 1 pile of 7.

When it is his turn, the player have to choose one of the following actions:

- Put one of a remaining warrior tile on an empty cell, with the number hidden;
- Put **one or two palisades** on the board, in an authorized border of a cell (one cannot put a palisade if it closes a territory of less than four cells);
- Use the race's power (only in advanced rules);
- Use reinforcement (only in advanced rules);
- Pass his turn.

Once a player passed his turn, he cannot play anymore for the rest of the game.

The game ends once every player has passed they turn. Every armies' strength is revealed and the gold of each territory is given to the player with the greatest army.

![End of a game](https://github.com/ThomasFerro/readmes/blob/master/posts/12-ddd-in-action-armadora-the-board-game/end_of_a_game.jpg)

In case of a tie, the players will compare their piles of gold, from highest to lowest.

Example: If the Elf have a pile of 6, a pile of 4 and a pile of 3 and the Orc have a pile of 6, one of 5 and one of 2, the Orc wins.

In a four-player game, the facing players can play as partners. We will not be implementing this feature in the first version.
