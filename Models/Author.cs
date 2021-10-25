using System.ComponentModel.DataAnnotations.Schema;

namespace estuary_backend.Models
{
    //TODO: after integrating with drops
    public class Author
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id { get; set; }

        public string DropsID {  get; set; }

        public bool Visible { get; set; }
    }
}
