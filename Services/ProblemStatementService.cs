using estuary_backend.Models;
using System.Linq;

namespace estuary_backend.Services
{
    public interface IProblemStatementService
    {
        bool DeleteProblemStatement(int id);
        ProblemStatement GetProblemStatementById(int id);

    }

    public class ProblemStatementService : IProblemStatementService
    {
        public bool DeleteProblemStatement(int id)
        {
            using var ctx = new EstuaryDbContext();
            var problemStatement = GetProblemStatementById(id);
            if (problemStatement is not null)
            {
                ctx.Contents.Remove(problemStatement);
                ctx.SaveChanges();
                return true;
            }
            return false;
        }

        public ProblemStatement GetProblemStatementById(int id)
        {
            using var ctx = new EstuaryDbContext();
            var problemStatement = ctx.Contents
                .Where(content => content.Id == id)
                .Where(content => content is ProblemStatement)
                .Select(content => content as ProblemStatement)
                .FirstOrDefault();

            return problemStatement;
        }
    }
}
