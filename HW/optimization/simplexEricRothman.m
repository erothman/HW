function [xsol,optimalobjective,basisfinal]=simplexEricRothman(A,b,c,BAS)

optimalobjective = 0;

c_tab = -transpose(c);

for initial = 1:length(BAS)
    
    col = BAS(initial);
    b(initial) = b(initial)/A(initial,col);
    A(initial,:) = A(initial,:)/A(initial,col);
    optimalobjective = optimalobjective - c_tab(1,col)*b(initial);
    c_tab(1,:) = c_tab(1,:) - c_tab(1,col)*A(initial,:);
    for rowChange = 1:length(BAS)
        if initial ~= rowChange
            b(rowChange) = b(rowChange) - A(rowChange, col)*b(initial);
            A(rowChange,:) = A(rowChange,:) - A(rowChange, col)*A(initial,:);
        end
    end
end
while max(c_tab) > 0
    if min(b) < 0
        error_msg = "The conditions contain a negative.\nSo the problem is infeasible.\n";
        error(error_msg);
    end
    [M,I] = max(c_tab);
    ratios = zeros(1,length(b));
    for row = 1:length(b)
        if A(row,I) > 0
            ratios(row) = b(row)/A(row,I);
        else
            ratios(row) = -1;
        end
    end
    if max(ratios) < 0
        error_msg = "The column only contains 0 and negatives.\nSo the problem is unbounded\n";
        error(error_msg);
        
    end
    for rowAgain = 1:length(ratios)
        if ratios(rowAgain) == -1
            ratios(rowAgain) = max(ratios) + 1;
        end
    end
    [MinRatio,minRow] = min(ratios);
    b(minRow) = b(minRow)/A(minRow,I);
    A(minRow,:) = A(minRow,:)/A(minRow,I);
    optimalobjective = optimalobjective - c_tab(1,I)*b(minRow);
    c_tab(1,:) = c_tab(1,:) - c_tab(1,I)*A(minRow,:);
    for rowChange = 1:length(BAS)
        if minRow ~= rowChange
            b(rowChange) = b(rowChange) - A(rowChange, I)*b(minRow);
            A(rowChange,:) = A(rowChange,:) - A(rowChange, I)*A(minRow,:);
        end
    end
end

xsol = zeros(1,length(c_tab));
basisfinal = zeros(1,length(b));

for finalRows = 1:length(b)
   for finalCols = 1:length(c_tab(1,:))
       if A(finalRows,finalCols) == 1 & c_tab(1,finalCols) == 0
           flag = 0;
           for row = 1:length(b)
               if row ~= finalRows & A(row,finalCols) == 1
                   flag = 1;
               end
           end
           if flag == 0
               xsol(finalCols) = b(finalRows);
               basisfinal(finalRows) = finalCols;
           end
       end
   end
end
end