library(ggplot2)

d <- read.table("final-aws.csv", header=T, sep=",")

#################
p = ggplot(d, aes(x = TotThd, y = SumRate))

p + facet_grid(Label ~ Test) + geom_boxplot(aes(group= cut_width(TotThd, 4)))

p.sel = ggplot(d[d$Test == "select-by-pk",], aes(x = TotThd, y = SumRate)) # + coord_trans(x="log2")

p.sel + facet_grid(. ~ Label) + geom_boxplot(aes(group= cut_width(TotThd, 4))) + 
xlab('Client Connections') + ylab('Operations / s')


dev.copy2pdf(file= "pk-select-kvm.pdf") 

####
p.sel = ggplot(d[d$Test == "insert-by-pk",], aes(x = TotThd, y = SumRate)) # + coord_trans(x="log2")

p.sel + facet_grid(. ~ Label) + geom_boxplot(aes(group= cut_width(TotThd, 4))) + 
xlab('Client Connections') + ylab('Operations / s')


dev.copy2pdf(file= "pk-insert-kvm.pdf") 

#################

p.sel = ggplot(d[d$Test == "update-by-pk",], aes(x = TotThd, y = SumRate)) # + coord_trans(x="log2")

p.sel + facet_grid(. ~ Label) + geom_boxplot(aes(group= cut_width(TotThd, 4))) + 
xlab('Client Connections') + ylab('Operations / s')


dev.copy2pdf(file= "pk-update-kvm.pdf") 

#################

p.sel = ggplot(d[d$Test == "delete-by-pk",], aes(x = TotThd, y = SumRate)) # + coord_trans(x="log2")

p.sel + facet_grid(. ~ Label) + geom_boxplot(aes(group= cut_width(TotThd, 4))) + 
xlab('Client Connections') + ylab('Operations / s')


dev.copy2pdf(file= "pk-delete-kvm.pdf") 

