/**
MAZE = [[0,0,0,0,0],
        [1,0,1,0,1],
        [0,0,0,0,0],
        [0,1,1,1,1],
        [0,0,0,0,9]]
The mazes are passed in as a two dimensional array. 0's represent spaces you may traverse. 1's represent walls. The program always starts in the 0,0 coordinate of the maze which is guaranteed to be a 0. The 9 represents the end of the maze and may be positioned anywhere. In this example maze it lives at 4,4. You may only travel up, down, left and right, no diagonal moves are allowed. The maze has already been passed through a maze validator so it is guaranteed not to contain any bad types or bad values. No letter 'a' or -1 will make it to your program.

The success condition is to return a list of coordinates that would let us retrace a valid path through the maze.
There may not be a valid path for every maze, in which case you should return a false value like an empty set or a None or a False.

If you complete this program we might start to talk about the shortest path through the maze or the shortest running time solution, but to start any path through the maze is good. We might also later talk about very very large mazes, not just 5x5 mazes, but milliion x million element mazes.
**/

package algo
