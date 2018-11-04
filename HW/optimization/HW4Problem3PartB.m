C = [1,7,-37,94,-9,76,-146]
b = [26;18;34]
A = [7,7,45,-1,3,-53,-68;9,-5,27,-115,7,-129,42;5,-3,63,-96,10,-109,86]
rightSide = [0;b]
leftSide = [1;0;0;0]
middle = [-C; A]
Tableau = [leftSide,middle,rightSide]

fprintf("This is the pretableau\n")
fprintf("The starting Basis we will work with is columns 1,2,7\n")
format short
Tableau(2,:) = Tableau(2,:)/(Tableau(2,2))
Tableau(1,:) = Tableau(1,:)-(Tableau(2,:)*Tableau(1,2))
Tableau(3,:) = Tableau(3,:)-(Tableau(2,:)*Tableau(3,2))
Tableau(4,:) = Tableau(4,:)-(Tableau(2,:)*Tableau(4,2))
fprintf("Reduce the first column to rref\n")
Tableau(3,:) = Tableau(3,:)/(Tableau(3,3))
Tableau(1,:) = Tableau(1,:)-(Tableau(3,:)*Tableau(1,3))
Tableau(2,:) = Tableau(2,:)-(Tableau(3,:)*Tableau(2,3))
Tableau(4,:) = Tableau(4,:)-(Tableau(3,:)*Tableau(4,3))
fprintf("Reduce the 2nd column to rref\n")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,8))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,8))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,8))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,8))
fprintf("Reduce the 7th column to rref\n")
fprintf("This is one of the answers I got in part a when I continued ")
fprintf("to pivot on the column that had a 0 for the change.\n")
fprintf("As I showed then, if we pivot on that column, which is the only ")
fprintf("pivotable column, it will just loop between a few different bfses.\n")
fprintf("So the final answer I will submit here is the one that I go first, ")
fprintf("the one from the initial basis.\n")
fprintf("So the min value for the objective function is -22 which occurs at ")
fprintf("the values of:\n")
fprintf("[x1    [2.8\n x2     4.8\n x3     0\n x4  =  0\n x5     0\n x6     0\n x7]    0.4]\n")
fprintf("\n")
fprintf("This is a different bfs to what I gave as the final answer in part a, ")
fprintf("but the objective function value is the same.\n")
fprintf("The reason that both answers are correct despite being different ")
fprintf("bfs is that when pivoted on the column with a change of 0, the ")
fprintf("obj func value will not change.\nSince the two answers pivot into ")
fprintf("eachother, both have a change of 0, without changing the obj func")
fprintf(" value, the bfs, while different points, are both equally valid solutions.")