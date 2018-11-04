function [xsol,optimalobjective,basisfinal]=simplextwoEricRothman(A,b,c)
c_tab = -transpose(c);

cA = zeros(1,length(c_tab));
cSlack = zeros(1, length(b));
cSlack(:) = cSlack(:) - 1;
cPhase1 = [cA,cSlack];
Aphase1 = A;
phase1Row1 = zeros(1,length(b));
phase1Row1(1) = 1;
BAS = zeros(1,length(b));
BAS(1) = 1 + length(c_tab);
IdentityPhase1 = [phase1Row1];
for phase1spot = 2:length(b)
    phase1Row = zeros(1,length(b));
    phase1Row(phase1spot) = 1;
    IdentityPhase1 = [IdentityPhase1;phase1Row];
    BAS(phase1spot) = phase1spot + length(c_tab);
end
Aphase1 = [Aphase1,IdentityPhase1];

optimalobjective = 0;

for initial = 1:length(BAS)
    
    col = BAS(initial);
    b(initial) = b(initial)/Aphase1(initial,col);
    Aphase1(initial,:) = Aphase1(initial,:)/Aphase1(initial,col);
    optimalobjective = optimalobjective - cPhase1(1,col)*b(initial);
    cPhase1(1,:) = cPhase1(1,:) - cPhase1(1,col)*Aphase1(initial,:);
    for rowChange = 1:length(BAS)
        if initial ~= rowChange
            b(rowChange) = b(rowChange) - Aphase1(rowChange, col)*b(initial);
            Aphase1(rowChange,:) = Aphase1(rowChange,:) - Aphase1(rowChange, col)*Aphase1(initial,:);
        end
    end
end
while max(cPhase1) > 0.0000001
    [M,I] = max(cPhase1);
    ratios = zeros(1,length(b));
    for row = 1:length(b)
        if Aphase1(row,I) > 0
            ratios(row) = b(row)/Aphase1(row,I);
        else
            ratios(row) = -1;
        end
    end
    if max(ratios) >= 0
        
    end
    for rowAgain = 1:length(ratios)
        if ratios(rowAgain) == -1
            ratios(rowAgain) = max(ratios) + 1;
        end
    end
    [MinRatio,minRow] = min(ratios);
    b(minRow) = b(minRow)/Aphase1(minRow,I);
    Aphase1(minRow,:) = Aphase1(minRow,:)/Aphase1(minRow,I);
    optimalobjective = optimalobjective - cPhase1(1,I)*b(minRow);
    cPhase1(1,:) = cPhase1(1,:) - cPhase1(1,I)*Aphase1(minRow,:);
    for rowChange = 1:length(BAS)
        if minRow ~= rowChange
            b(rowChange) = b(rowChange) - Aphase1(rowChange, I)*b(minRow);
            Aphase1(rowChange,:) = Aphase1(rowChange,:) - Aphase1(rowChange, I)*Aphase1(minRow,:);
        end
    end
end

if optimalobjective > 0.000001 | optimalobjective < -0.000001
    err_msg = "The optimal obj. value to phase 1 is not zero.\nSo the problem is infeasible.\n";
    error(err_msg);
end

phase1xsol = zeros(1,length(cPhase1));
startingBasis = zeros(1,length(b));

for finalRows = 1:length(b)
   for finalCols = 1:length(cPhase1(1,:))
       if Aphase1(finalRows,finalCols) == 1 & cPhase1(1,finalCols) == 0
           flag = 0;
           for row = 1:length(b)
               if row ~= finalRows & Aphase1(row,finalCols) == 1
                   flag = 1;
               end
           end
           if flag == 0
               phase1xsol(finalCols) = b(finalRows);
               startingBasis(finalRows) = finalCols;
           end
       end
   end
end

for cols = 1:length(c_tab)
    for rows = 1:length(b)
        A(rows,cols) = Aphase1(rows,cols);
    end
end

[xsol, optimalobjective, basisfinal] = simplexEricRothman(A, b, c, startingBasis);
end