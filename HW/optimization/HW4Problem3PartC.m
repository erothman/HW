C = [1,7,-37,94,-9,76,-146]
b = [26;18;34]
A = [7,7,45,-1,3,-53,-68;9,-5,27,-115,7,-129,42;5,-3,63,-96,10,-109,86]
rightSide = [0;b]
leftSide = [1;0;0;0]
middle = [-C; A]
Tableau = [leftSide,middle,rightSide]

fprintf("This is the pretableau\n")
fprintf("The starting Basis we will work with is columns 1,3,6\n")
format short
Tableau(2,:) = Tableau(2,:)/(Tableau(2,2))
Tableau(1,:) = Tableau(1,:)-(Tableau(2,:)*Tableau(1,2))
Tableau(3,:) = Tableau(3,:)-(Tableau(2,:)*Tableau(3,2))
Tableau(4,:) = Tableau(4,:)-(Tableau(2,:)*Tableau(4,2))
fprintf("Reduce the first column to rref\n")
Tableau(3,:) = Tableau(3,:)/(Tableau(3,4))
Tableau(1,:) = Tableau(1,:)-(Tableau(3,:)*Tableau(1,4))
Tableau(2,:) = Tableau(2,:)-(Tableau(3,:)*Tableau(2,4))
Tableau(4,:) = Tableau(4,:)-(Tableau(3,:)*Tableau(4,4))
fprintf("Reduce the 3rd column to rref\n")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,7))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,7))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,7))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,7))
fprintf("Reduce the 6th column to rref\n")
fprintf("The Column that would cause the largest change is column 4:")
fprintf(" it is the one with the largest number in the top row.\n")
fprintf("So we will pivot based on that column.\n")
fprintf("We will pivot based on the row with the lowest ratio, which in ")
fprintf("this case is the 4th row since its ration is 0.\n")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,5))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,5))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,5))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,5))
fprintf("Now the column with the largest change is column 2.\n")
fprintf("Once again, since the the ratio on row 4 is 0 since the b value is 0, ")
fprintf("We will pivot on column 2 at the 4th row.")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,3))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,3))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,3))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,3))
fprintf("Now the column with the largest change is column 7.\n")
fprintf("Unlike the previous times though, the value in column 7 on the ")
fprintf("4th row is negative, so we will ignore the ratio of the 4th row.\n")
fprintf("Infact the only pivotable row that does not have a negative is row 3.\n")
fprintf("So we will pivot about column 7, row 3.")
Tableau(3,:) = Tableau(3,:)/(Tableau(3,8))
Tableau(1,:) = Tableau(1,:)-(Tableau(3,:)*Tableau(1,8))
Tableau(2,:) = Tableau(2,:)-(Tableau(3,:)*Tableau(2,8))
Tableau(4,:) = Tableau(4,:)-(Tableau(3,:)*Tableau(4,8))
fprintf("Once again we arrive at the answer to part b.\nSo if we pivot on ")
fprintf("the column with a change of 0, we will just loop through some answers.\n")
fprintf("So this is where we end this algorithm run.\n")
fprintf("So the min value for the objective function is -22 which occurs at ")
fprintf("the values of:\n")
fprintf("[x1    [2.8\n x2     4.8\n x3     0\n x4  =  0\n x5     0\n x6     0\n x7]    0.4]\n")
fprintf("\n")
fprintf("Interestingly with this run of the algorithm, we saw what happens ")
fprintf("when we used an initial basis that was not one of the answers.\n")
fprintf("The algorithm stalled thanks to the degeneracy in the 4th row.\n")
fprintf("The stalling ended only when the algorithm moved onwards to one of ")
fprintf("the final answers.\nAnytime the algorithm moved between basis that ")
fprintf("were not correct answers, there was no change in the obj. func. value.\n")
fprintf("Its like there are different levels for the bases, where each level ")
fprintf("contains the bases that compute the obj. func. to a certain value.\n")
fprintf("So there is the -18 level, which contains bases made of columns ")
fprintf("(1,3,7),(1,3,4),(1,2,3).\nThen there is level -22, which is the ")
fprintf("lowest level and contains the bases of columns (1,2,5),(1,2,7),(2,5,7).\n")
