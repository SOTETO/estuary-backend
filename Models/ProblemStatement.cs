using estuary_backend.DTOs;
using System.Linq;

namespace estuary_backend.Models
{
    public class ProblemStatement : Content
    {
        public string Iam { get; set; }

        public string Iwant { get; set; }

        public string But { get; set; }

        public string Because { get; set; }

        public string Feel { get; set; }

        public ProblemStatementDTO toDTO()
        {
            return new ProblemStatementDTO
            {
                id = Id,
                likes = Likes.Count,
                iAm = Iam,
                iWant = Iwant,
                but = But,
                because = Because,
                feel = Feel,
                linked = RelatedContent.Select(related => related.toDTO()).ToArray(),
            };
        }
    }
}
